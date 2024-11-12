package models

type WlanPerformanceReq struct {
	BssIds          []string `json:"bssIds,omitempty"`
	WlanOrderCriter string   `json:"wlanOrderCriter,omitempty"`
	Order           string   `json:"order,omitempty"`
	SsIds           []string `json:"ssIds,omitempty"`
}
