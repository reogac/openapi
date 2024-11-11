package models

type Tai struct {
	Tac    string `json:"tac"`
	Nid    string `json:"nid,omitempty"`
	PlmnId PlmnId `json:"plmnId"`
}
