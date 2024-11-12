package models

type WlanPerformanceReq struct {
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty"`
}
