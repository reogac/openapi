package models

type MbsServiceInfo struct {
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
	MbsSdfResPrio  ReservPriority          `json:"mbsSdfResPrio,omitempty"`
	AfAppId        string                  `json:"afAppId,omitempty"`
}
