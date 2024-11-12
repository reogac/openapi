package models

type MbsServiceInfo struct {
	AfAppId        string                  `json:"afAppId,omitempty"`
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
	MbsSdfResPrio  ReservPriority          `json:"mbsSdfResPrio,omitempty"`
}
