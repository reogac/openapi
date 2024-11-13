package models

type Ncgi struct {
	NrCellId string `json:"nrCellId"`
	Nid      string `json:"nid,omitempty"`
	PlmnId   PlmnId `json:"plmnId"`
}
