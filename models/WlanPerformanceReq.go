package models
type WlanPerformanceReq struct {
	 BssIds	[]string	`json:"bssIds,omitempty"`
	 WlanOrderCriter	WlanOrderingCriterion	`json:"wlanOrderCriter,omitempty"`
	 Order	MatchingDirection	`json:"order,omitempty"`
	 SsIds	[]string	`json:"ssIds,omitempty"`
}
