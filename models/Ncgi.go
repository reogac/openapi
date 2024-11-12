package models

type Ncgi struct {
	Nid      string `json:"nid,omitempty"`
	PlmnId   PlmnId `json:"plmnId"`
	NrCellId string `json:"nrCellId"`
}
