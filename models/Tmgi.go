package models

type Tmgi struct {
	MbsServiceId string `json:"mbsServiceId"`
	PlmnId       PlmnId `json:"plmnId"`
}
