package openrtb

import (
	"bytes"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/minio/simdjson-go"
)

//easyjson:json
type Imp struct {
	Bidfloor       float64    `json:"bidfloor"`
	Secure         int        `json:"secure"`
	BidfloorCur    string     `json:"bidfloorcur"`
	DisplayManager string     `json:"displaymanager"`
	ID             string     `json:"id"`
	Banner         *ImpBanner `json:"banner"`
	Video          *ImpVideo  `json:"video"`
}

//easyjson:json
type ImpBanner struct {
	H      int64        `json:"h"`
	W      int64        `json:"w"`
	Pos    int64        `json:"pos"`
	API    []int64      `json:"api"`
	Expdir []int64      `json:"expdir"`
	Format []*ImpFormat `json:"format"`
}

//easyjson:json
type ImpFormat struct {
	H int64 `json:"h"`
	W int64 `json:"w"`
}

//easyjson:json
type ImpVideo struct {
	MinDuration int64   `json:"minduration"`
	MaxDuration int64   `json:"maxduration"`
	Protocol    int64   `json:"protocol"`
	W           int64   `json:"w"`
	H           int64   `json:"h"`
	StartDelay  int64   `json:"startdelay"`
	Placement   int64   `json:"placement"`
	Protocols   []int64 `json:"protocols"`
	Battr       []int64 `json:"battr"`
	Mimes       []int64 `json:"mimes"`
}

const (
	fieldBidfloor fieldIdx = iota
	fieldSecure
	fieldBannerW
	fieldBannerH
	fieldBannerFormat
	fieldVideoMaxDur
	fieldVideoMinDur
	fieldVideoMimes
	fieldVideoProtocols
	fieldVideoBattr
	fieldVideoW
	fieldVideoH
)

var (
	impFields = []rtbFieldDef{
		{fieldBidfloor, []string{"bidfloor"}},
		{fieldSecure, []string{"secure"}},
		{fieldBannerW, []string{"banner", "w"}},
		{fieldBannerH, []string{"banner", "h"}},
		{fieldBannerH, []string{"banner", "h"}},
		{fieldBannerFormat, []string{"banner", "format"}},
		{fieldVideoMaxDur, []string{"video", "maxduration"}},
		{fieldVideoMinDur, []string{"video", "minduration"}},
		{fieldVideoMimes, []string{"video", "mimes"}},
		{fieldVideoProtocols, []string{"video", "protocols"}},
		{fieldVideoBattr, []string{"video", "battr"}},
		{fieldVideoW, []string{"video", "w"}},
		{fieldVideoH, []string{"video", "h"}},
	}

	impPaths = rtbBuildPaths(impFields)
)

func (i *Imp) setField(idx int, value []byte, _ jsonparser.ValueType, _ error) {
	switch fieldIdx(idx) {
	case fieldBidfloor:
		i.Bidfloor, _ = strconv.ParseFloat(string(value), 64)
	case fieldSecure:
		i.Secure, _ = strconv.Atoi(string(value))
	case fieldBannerW:
		i.Banner.W, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldBannerH:
		i.Banner.H, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldBannerFormat:
		i.Banner.Format = []*ImpFormat{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			iformat := &ImpFormat{}
			w, _, _, werr := jsonparser.Get(arrdata, "w")
			h, _, _, herr := jsonparser.Get(arrdata, "h")
			if werr != nil || herr != nil {
				return
			}
			iformat.W, _ = strconv.ParseInt(string(w), 10, 64)
			iformat.H, _ = strconv.ParseInt(string(h), 10, 64)
			i.Banner.Format = append(i.Banner.Format, iformat)
		})
	case fieldVideoMaxDur:
		i.Video.MaxDuration, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldVideoMinDur:
		i.Video.MinDuration, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldVideoMimes:
		i.Video.Mimes = []int64{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			ival, err := strconv.ParseInt(string(arrdata), 10, 64)
			if err != nil {
				return
			}
			i.Video.Mimes = append(i.Video.Mimes, ival)
		})
	case fieldVideoProtocols:
		i.Video.Protocols = []int64{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			ival, err := strconv.ParseInt(string(arrdata), 10, 64)
			if err != nil {
				return
			}
			i.Video.Protocols = append(i.Video.Protocols, ival)
		})
	case fieldVideoBattr:
		i.Video.Battr = []int64{}
		jsonparser.ArrayEach(value, func(arrdata []byte, dataType jsonparser.ValueType, offset int, err error) {
			ival, err := strconv.ParseInt(string(arrdata), 10, 64)
			if err != nil {
				return
			}
			i.Video.Battr = append(i.Video.Battr, ival)
		})
	case fieldVideoW:
		i.Video.W, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldVideoH:
		i.Video.H, _ = strconv.ParseInt(string(value), 10, 64)
	}
}

func (i *Imp) UnmarshalJSONReq(b []byte) error {
	i.Banner = &ImpBanner{}
	i.Video = &ImpVideo{}
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		i.setField(idx, value, vt, err)
	}, impPaths...)

	return nil
}

func parseImps(arr *simdjson.Array, tmp *simdjson.Iter, obj *simdjson.Object) ([]*Imp, error) {
	iter := arr.Iter()
	imps := []*Imp{}
	for {
		t := iter.Advance()
		if t == simdjson.TypeNone {
			return imps, nil
		}

		if t != simdjson.TypeObject {
			return nil, ErrExpectedObject
		}

		if _, err := iter.Object(obj); err != nil {
			return nil, err
		}

		imp := &Imp{}

		for {
			keyName, t, err := obj.NextElementBytes(tmp)
			if err != nil {
				return nil, err
			}

			if t == simdjson.TypeNone {
				break
			}

			switch t {
			case simdjson.TypeFloat:
				f, err := tmp.Float()
				if err != nil {
					return nil, err
				}
				switch {
				case bytes.Compare(keyName, bidfloorKey) == 0:
					imp.Bidfloor = f
				}
			}
		}

		imps = append(imps, imp)
	}
}
