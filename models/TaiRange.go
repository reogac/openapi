package models

type TaiRange struct {
	TacRangeList []TacRange `json:"tacRangeList"`
	Nid          string     `json:"nid,omitempty"`
	PlmnId       PlmnId     `json:"plmnId"`
}
