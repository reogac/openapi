package models

type WlanPerformanceReq struct {
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty"`
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
}
