package models

type WlanPerformanceReq struct {
	SsIds           []string `json:"ssIds,omitempty"`
	BssIds          []string `json:"bssIds,omitempty"`
	WlanOrderCriter string   `json:"wlanOrderCriter,omitempty"`
	Order           string   `json:"order,omitempty"`
}
