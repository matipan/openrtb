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
	IP             []byte     `json:"ip"`
	UA             []byte     `json:"ua"`
	Language       []byte     `json:"language"`
	OS             []byte     `json:"os"`
	OSV            []byte     `json:"osv"`
	IFA            []byte     `json:"ifa"`
	HWV            []byte     `json:"hwv"`
	Model          []byte     `json:"model"`
	DNT            []byte     `json:"dnt"`
	H              int64      `json:"h"`
	W              int64      `json:"w"`
	ConnectionType int64      `json:"connectionType"`
	DeviceType     int64      `json:"deviceType"`
	Ext            *DeviceExt `json:"ext"`
}

//easyjson:json
type DeviceExt struct {
	IFV []byte `json:"ifv"`
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
				d.IP = b
			case bytes.Compare(name, uaKey) == 0:
				d.UA = b
			case bytes.Compare(name, osKey) == 0:
				d.OS = b
			case bytes.Compare(name, osvKey) == 0:
				d.OSV = b
			case bytes.Compare(name, ifaKey) == 0:
				d.IFA = b
			case bytes.Compare(name, hwvKey) == 0:
				d.HWV = b
			case bytes.Compare(name, modelKey) == 0:
				d.Model = b
			case bytes.Compare(name, dntKey) == 0:
				d.DNT = b
			case bytes.Compare(name, langKey) == 0:
				d.Language = b
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
		data.Language = value
	case fieldDeviceIp: // ip
		data.IP = value
	case fieldDeviceUa: // ua
		data.UA = value
	case fieldDeviceOs: // os
		data.OS = value
	case fieldDeviceOsv: // osv
		data.OSV = value
	case fieldDeviceIfa: // ifa
		data.IFA = value
	case fieldDeviceHwv: // hwv
		data.HWV = value
	case fieldDeviceModel: //model
		data.Model = value
	case fieldDeviceDnt: // dnt
		data.DNT = value
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
