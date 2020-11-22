package openrtb

import (
	"github.com/buger/jsonparser"
)

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
		a.Name = value
	case fieldAppPubId:
		a.Publisher.ID = value
	case fieldAppBundle:
		a.Bundle = value
	case fieldAppLanguage:
		a.Content.Language = value
	case fieldAppId:
		a.ID = value
	case fieldExtDevUserId:
		a.Ext.Devuserid = value
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
