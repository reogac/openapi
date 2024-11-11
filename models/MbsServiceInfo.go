package models

type MbsServiceInfo struct {
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
	MbsSdfResPrio  string                  `json:"mbsSdfResPrio,omitempty"`
	AfAppId        string                  `json:"afAppId,omitempty"`
}
