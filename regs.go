package openrtb

//easyjson:json
type Regs struct {
	COPPA int64    `json:"coppa"`
	Ext   *RegsExt `json:"ext"`
}

//easyjson:json
type RegsExt struct {
	GDPR []byte `json:"gdpr"`
	CCPA []byte `json:"us_privacy"`
}
