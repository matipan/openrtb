package openrtb

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/minio/simdjson-go"
)

var (
	dKey        = []byte("device")
	impKey      = []byte("imp")
	hKey        = []byte("h")
	wKey        = []byte("w")
	bidfloorKey = []byte("bidfloor")
	extKey      = []byte("ext")
)

var (
	ErrExpectedInt    = errors.New("expected integer")
	ErrExpectedObject = errors.New("expected object")
)

//easyjson:json
type Request struct {
	At     int64    `json:"at"`
	Tmax   int64    `json:"tmax"`
	ID     string   `json:"id"`
	BCat   []string `json:"bcat"`
	BAdv   []string `json:"badv"`
	BApp   []string `json:"bapp"`
	Imps   []*Imp   `json:"imp"`
	Device *Device  `json:"device"`
	App    *App     `json:"app"`
	Regs   *Regs    `json:"regs"`
}

const (
	fieldDevice fieldIdx = iota
	fieldImp
	fieldApp
	fieldId
	fieldAt
	fieldBcat
	fieldBadv
	fieldBapp
	fieldRegs
)

var (
	reqFields = []rtbFieldDef{
		{fieldDevice, []string{"device"}},
		{fieldImp, []string{"imp"}},
		{fieldApp, []string{"app"}},
		{fieldId, []string{"id"}},
		{fieldAt, []string{"at"}},
		{fieldBcat, []string{"bcat"}},
		{fieldBadv, []string{"badv"}},
		{fieldBapp, []string{"bapp"}},
		{fieldRegs, []string{"regs"}},
	}

	reqPaths = rtbBuildPaths(reqFields)
)

func (data *Request) setField(idx int, value []byte, _ jsonparser.ValueType, _ error) {
	switch fieldIdx(idx) {
	case fieldDevice:
		data.Device = &Device{}
		data.Device.UnmarshalJSONReq(value)
	case fieldImp:
		data.Imps = []*Imp{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			imp := &Imp{}
			if err := imp.UnmarshalJSONReq(value); err != nil {
				return
			}
			data.Imps = append(data.Imps, imp)
		})
	case fieldApp:
		data.App = &App{}
		data.App.UnmarshalJSONReq(value)
	case fieldId:
		data.ID = string(value)
	case fieldAt:
		data.At, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldBcat:
		data.BCat = []string{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			data.BCat = append(data.BCat, string(value))
		})
	case fieldBadv:
		data.BAdv = []string{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			data.BAdv = append(data.BAdv, string(value))
		})
	case fieldBapp:
		data.BApp = []string{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			data.BApp = append(data.BApp, string(value))
		})
	case fieldRegs:
		data.Regs = &Regs{}
		data.Regs.UnmarshalJSONReq(value)
	}
}

func (r *Request) UnmarshalJSONReq(b []byte) error {
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		r.setField(idx, value, vt, err)
	}, reqPaths...)

	return nil
}

func (r *Request) UnmarshalJSONSimd(b []byte) error {
	parsed, err := simdjson.Parse(b, nil)
	if err != nil {
		return err
	}

	var (
		iter = parsed.Iter()
		obj  = &simdjson.Object{}
		tmp  = &simdjson.Iter{}
	)

	for {
		typ := iter.Advance()

		switch typ {
		case simdjson.TypeRoot:
			if typ, tmp, err = iter.Root(tmp); err != nil {
				return err
			}

			switch typ {
			case simdjson.TypeObject:
				if obj, err = tmp.Object(obj); err != nil {
					return err
				}
				return r.parse(tmp, obj)
			}
		default:
			return nil
		}
	}
}

func (r *Request) parse(tmp *simdjson.Iter, obj *simdjson.Object) error {
	arr := &simdjson.Array{}
	for {
		name, t, err := obj.NextElementBytes(tmp)
		if err != nil {
			return err
		}

		if t == simdjson.TypeNone {
			break
		}

		switch t {
		case simdjson.TypeObject:
			if err := r.parseObject(name, tmp); err != nil {
				return err
			}
		case simdjson.TypeArray:
			if _, err := tmp.Array(arr); err != nil {
				return err
			}
			if err := r.parseArray(name, tmp, arr); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Request) parseArray(name []byte, tmp *simdjson.Iter, arr *simdjson.Array) (err error) {
	obj := &simdjson.Object{}
	switch {
	case bytes.Compare(name, impKey) == 0:
		r.Imps, err = parseImps(arr, tmp, obj)
		return err
	}

	return nil
}

func (r *Request) parseObject(name []byte, iter *simdjson.Iter) error {
	obj := &simdjson.Object{}
	switch {
	case bytes.Compare(name, dKey) == 0:
		if _, err := iter.Object(obj); err != nil {
			return err
		}
		r.Device = &Device{}
		return r.Device.parse(iter, obj)
	}
	return nil
}
