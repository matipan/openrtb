package openrtb

import (
	"strconv"

	"github.com/buger/jsonparser"
)

//easyjson:json
type Regs struct {
	COPPA int64    `json:"coppa"`
	Ext   *RegsExt `json:"ext"`
}

//easyjson:json
type RegsExt struct {
	GDPR string `json:"gdpr"`
	CCPA string `json:"us_privacy"`
}

const (
	fieldRegsCoppa fieldIdx = iota
	fieldRegsExtGdpr
	fieldRegsCCPA
)

var (
	regsFields = []rtbFieldDef{
		{fieldRegsCoppa, []string{"regs", "coppa"}},
		{fieldRegsExtGdpr, []string{"regs", "ext", "gdpr"}},
		{fieldRegsCCPA, []string{"regs", "ext", "us_privacy"}},
	}

	regsPaths = rtbBuildPaths(regsFields)
)

func (r *Regs) setField(idx int, value []byte, _ jsonparser.ValueType, _ error) {
	switch fieldIdx(idx) {
	case fieldRegsCoppa:
		r.COPPA, _ = strconv.ParseInt(string(value), 10, 64)
	case fieldRegsExtGdpr:
		r.Ext.GDPR = string(value)
	case fieldRegsCCPA:
		r.Ext.CCPA = string(value)
	}
}

func (r *Regs) UnmarshalJSONReq(b []byte) error {
	r.Ext = &RegsExt{}
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		r.setField(idx, value, vt, err)
	}, regsPaths...)

	return nil
}
