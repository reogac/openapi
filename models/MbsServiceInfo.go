package models

type MbsServiceInfo struct {
	MbsSdfResPrio  ReservPriority          `json:"mbsSdfResPrio,omitempty"`
	AfAppId        string                  `json:"afAppId,omitempty"`
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
}
