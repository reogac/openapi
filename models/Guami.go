package models

type Guami struct {
	AmfId  string    `json:"amfId"`
	PlmnId PlmnIdNid `json:"plmnId"`
}
