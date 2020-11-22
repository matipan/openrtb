package openrtb

//easyjson:json
type Request struct {
	At     int64    `json:"at"`
	Tmax   int64    `json:"tmax"`
	ID     []byte   `json:"id"`
	BCat   [][]byte `json:"bcat"`
	BAdv   [][]byte `json:"badv"`
	BApp   [][]byte `json:"bapp"`
	Imps   []*Imp   `json:"imp"`
	Device *Device  `json:"device"`
	App    *App     `json:"app"`
	Regs   *Regs    `json:"regs"`
}
