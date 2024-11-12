package models

type Tai struct {
	Nid    string `json:"nid,omitempty"`
	PlmnId PlmnId `json:"plmnId"`
	Tac    string `json:"tac"`
}
