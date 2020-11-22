package openrtb

import (
	"bytes"

	"github.com/buger/jsonparser"
	"github.com/minio/simdjson-go"
)

var (
	ipKey    = []byte("ip")
	uaKey    = []byte("ua")
	osKey    = []byte("os")
	osvKey   = []byte("osv")
	ifaKey   = []byte("ifa")
	ctKey    = []byte("connectionType")
	hwvKey   = []byte("hwv")
	modelKey = []byte("model")
	dntKey   = []byte("dnt")
	langKey  = []byte("language")
	dtKey    = []byte("deviceType")
	ifvKey   = []byte("ifv")
	latKey   = []byte("lat")
	lonKey   = []byte("lon")
)

//easyjson:json
type Device struct {
	IP             string     `json:"ip"`
	UA             string     `json:"ua"`
	Language       string     `json:"language"`
	OS             string     `json:"os"`
	OSV            string     `json:"osv"`
	IFA            string     `json:"ifa"`
	HWV            string     `json:"hwv"`
	Model          string     `json:"model"`
	DNT            string     `json:"dnt"`
	H              int64      `json:"h"`
	W              int64      `json:"w"`
	ConnectionType int64      `json:"connectionType"`
	DeviceType     int64      `json:"deviceType"`
	Ext            *DeviceExt `json:"ext"`
}

//easyjson:json
type DeviceExt struct {
	IFV string `json:"ifv"`
}

func (d *Device) UnmarshalJSONSimd(b []byte) error {
	parsed, err := simdjson.Parse(b, nil)
	if err != nil {
		return err
	}
	iter := parsed.Iter()

	typ := iter.Advance()
	typ, tmp, err := iter.Root(nil)
	if err != nil {
		return err
	}

	if typ != simdjson.TypeObject {
		return ErrExpectedObject
	}

	obj, err := tmp.Object(nil)
	if err != nil {
		return err
	}

	return d.parse(tmp, obj)
}

func (d *Device) parse(iter *simdjson.Iter, obj *simdjson.Object) error {
	for {
		name, t, err := obj.NextElementBytes(iter)
		if err != nil {
			return err
		}

		if t == simdjson.TypeNone {
			return nil
		}

		switch t {
		case simdjson.TypeInt:
			n, err := iter.Int()
			if err != nil {
				return err
			}
			switch {
			case bytes.Compare(name, hKey) == 0:
				d.H = n
			case bytes.Compare(name, wKey) == 0:
				d.W = n
			case bytes.Compare(name, dtKey) == 0:
				d.DeviceType = n
			case bytes.Compare(name, ctKey) == 0:
				d.ConnectionType = n
			}
		case simdjson.TypeString:
			b, err := iter.StringBytes()
			if err != nil {
				return err
			}
			switch {
			case bytes.Compare(name, ipKey) == 0:
				d.IP = string(b)
			case bytes.Compare(name, uaKey) == 0:
				d.UA = string(b)
			case bytes.Compare(name, osKey) == 0:
				d.OS = string(b)
			case bytes.Compare(name, osvKey) == 0:
				d.OSV = string(b)
			case bytes.Compare(name, ifaKey) == 0:
				d.IFA = string(b)
			case bytes.Compare(name, hwvKey) == 0:
				d.HWV = string(b)
			case bytes.Compare(name, modelKey) == 0:
				d.Model = string(b)
			case bytes.Compare(name, dntKey) == 0:
				d.DNT = string(b)
			case bytes.Compare(name, langKey) == 0:
				d.Language = string(b)
			}
		}
	}
}

type fieldIdx int

const (
	fieldDeviceLanguage fieldIdx = iota
	fieldDeviceIp
	fieldDeviceUa
	fieldDeviceOs
	fieldDeviceOsv
	fieldDeviceIfa
	fieldDeviceConnectionType
	fieldDeviceHwv
	fieldDeviceModel
	fieldDeviceGeoLat
	fieldDeviceGeoLon
	fieldDeviceType
	fieldDeviceDnt
	fieldDeviceExtIfv
)

type rtbFieldDef struct {
	idx  fieldIdx
	path []string
}

var deviceFields = []rtbFieldDef{
	{fieldDeviceLanguage, []string{"language"}},
	{fieldDeviceIp, []string{"ip"}},
	{fieldDeviceUa, []string{"ua"}},
	{fieldDeviceOs, []string{"os"}},
	{fieldDeviceOsv, []string{"osv"}},
	{fieldDeviceIfa, []string{"ifa"}},
	{fieldDeviceConnectionType, []string{"connectiontype"}},
	{fieldDeviceHwv, []string{"hwv"}},
	{fieldDeviceModel, []string{"model"}},
	{fieldDeviceGeoLat, []string{"geo", "lat"}},
	{fieldDeviceGeoLon, []string{"geo", "lon"}},
	{fieldDeviceType, []string{"devicetype"}},
	{fieldDeviceDnt, []string{"dnt"}},
	{fieldDeviceExtIfv, []string{"ext", "ifv"}},
}

func (data *Device) setField(idx int, value []byte, _ jsonparser.ValueType, _ error) {
	switch fieldIdx(idx) {
	case fieldDeviceLanguage: // language
		data.Language = string(value)
	case fieldDeviceIp: // ip
		data.IP = string(value)
	case fieldDeviceUa: // ua
		data.UA = string(value)
	case fieldDeviceOs: // os
		data.OS = string(value)
	case fieldDeviceOsv: // osv
		data.OSV = string(value)
	case fieldDeviceIfa: // ifa
		data.IFA = string(value)
	case fieldDeviceHwv: // hwv
		data.HWV = string(value)
	case fieldDeviceModel: //model
		data.Model = string(value)
	case fieldDeviceDnt: // dnt
		data.DNT = string(value)
	}
}

// build the path array we pass to the parser...
func rtbBuildPaths(fields []rtbFieldDef) [][]string {
	ret := make([][]string, 0, 10)
	for _, f := range fields {
		ret = append(ret, f.path)
	}
	return ret
}

// define fields to decode
var (
	devicePaths = rtbBuildPaths(deviceFields)
)

func (d *Device) UnmarshalJSONReq(b []byte) error {
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		d.setField(idx, value, vt, err)
	}, devicePaths...)
	return nil
}
