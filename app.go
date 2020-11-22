package openrtb

//easyjson:json
type App struct {
	Name      []byte      `json:"name"`
	Bundle    []byte      `json:"bundle"`
	ID        []byte      `json:"id"`
	Publisher *Publisher  `json:"publisher"`
	Content   *AppContent `json:"content"`
	Ext       *AppExt     `json:"ext"`
}

//easyjson:json
type Publisher struct {
	ID     []byte `json:"id"`
	Name   []byte `json:"name"`
	CAT    []byte `json:"cat"`
	Domain []byte `json:"domain"`
}

//easyjson:json
type AppContent struct {
	ID       []byte `json:"id"`
	Title    []byte `json:"title"`
	Language []byte `json:"language"`
}

//easyjson:json
type AppExt struct {
	Devuserid []byte `json:"devuserid"`
}
