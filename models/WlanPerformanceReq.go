package models

type WlanPerformanceReq struct {
	Order           MatchingDirection     `json:"order,omitempty"`
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
}
