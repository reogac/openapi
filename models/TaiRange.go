package models

type TaiRange struct {
	Nid          string     `json:"nid,omitempty"`
	PlmnId       PlmnId     `json:"plmnId"`
	TacRangeList []TacRange `json:"tacRangeList"`
}
