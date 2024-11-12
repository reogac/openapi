package models
type MbsServiceInfo struct {
	 MbsMediaComps	map[string]MbsMediaComp	`json:"mbsMediaComps"`
	 MbsSdfResPrio	string	`json:"mbsSdfResPrio,omitempty"`
	 AfAppId	string	`json:"afAppId,omitempty"`
	 MbsSessionAmbr	string	`json:"mbsSessionAmbr,omitempty"`
}
