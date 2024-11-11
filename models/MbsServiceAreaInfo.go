package models

type MbsServiceAreaInfo struct {
	AreaSessionId  int            `json:"areaSessionId"`
	MbsServiceArea MbsServiceArea `json:"mbsServiceArea"`
}
