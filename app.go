package openrtb

import (
	"github.com/buger/jsonparser"
)

//easyjson:json
type App struct {
	Name      string      `json:"name"`
	Bundle    string      `json:"bundle"`
	ID        string      `json:"id"`
	Publisher *Publisher  `json:"publisher"`
	Content   *AppContent `json:"content"`
	Ext       *AppExt     `json:"ext"`
}

//easyjson:json
type Publisher struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	CAT    string `json:"cat"`
	Domain string `json:"domain"`
}

//easyjson:json
type AppContent struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Language string `json:"language"`
}

//easyjson:json
type AppExt struct {
	Devuserid string `json:"devuserid"`
}

const (
	fieldAppName fieldIdx = iota
	fieldAppPubId
	fieldAppBundle
	fieldAppLanguage
	fieldAppId
	fieldExtDevUserId
)

var (
	appFields = []rtbFieldDef{
		{fieldAppName, []string{"name"}},
		{fieldAppPubId, []string{"publisher", "id"}},
		{fieldAppBundle, []string{"bundle"}},
		{fieldAppLanguage, []string{"content", "language"}},
		{fieldAppId, []string{"id"}},
		{fieldExtDevUserId, []string{"ext", "devuserid"}},
	}

	appPaths = rtbBuildPaths(appFields)
)

func (a *App) setField(idx int, value []byte, _ jsonparser.ValueType, _ error) {
	switch fieldIdx(idx) {
	case fieldAppName:
		a.Name = string(value)
	case fieldAppPubId:
		a.Publisher.ID = string(value)
	case fieldAppBundle:
		a.Bundle = string(value)
	case fieldAppLanguage:
		a.Content.Language = string(value)
	case fieldAppId:
		a.ID = string(value)
	case fieldExtDevUserId:
		a.Ext.Devuserid = string(value)
	}
}

func (a *App) UnmarshalJSONReq(b []byte) error {
	a.Publisher = &Publisher{}
	a.Content = &AppContent{}
	a.Ext = &AppExt{}
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		a.setField(idx, value, vt, err)
	}, appPaths...)

	return nil
}
