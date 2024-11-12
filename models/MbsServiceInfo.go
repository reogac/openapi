package models

type MbsServiceInfo struct {
	MbsMediaComps  map[string]MbsMediaComp `json:"mbsMediaComps"`
	MbsSdfResPrio  ReservPriority          `json:"mbsSdfResPrio,omitempty"`
	AfAppId        string                  `json:"afAppId,omitempty"`
	MbsSessionAmbr string                  `json:"mbsSessionAmbr,omitempty"`
}
