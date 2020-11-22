package openrtb

//easyjson:json
type Imp struct {
	Bidfloor       float64    `json:"bidfloor"`
	Secure         int        `json:"secure"`
	BidfloorCur    []byte     `json:"bidfloorcur"`
	DisplayManager []byte     `json:"displaymanager"`
	ID             []byte     `json:"id"`
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
