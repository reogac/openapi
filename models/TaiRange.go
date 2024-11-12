package models

type TaiRange struct {
	PlmnId       PlmnId     `json:"plmnId"`
	TacRangeList []TacRange `json:"tacRangeList"`
	Nid          string     `json:"nid,omitempty"`
}
