package openrtb

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
