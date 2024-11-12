package models

type Ecgi struct {
	EutraCellId string `json:"eutraCellId"`
	Nid         string `json:"nid,omitempty"`
	PlmnId      PlmnId `json:"plmnId"`
}
