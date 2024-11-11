package models

type CellGlobalId struct {
	Lac    string `json:"lac"`
	CellId string `json:"cellId"`
	PlmnId PlmnId `json:"plmnId"`
}
