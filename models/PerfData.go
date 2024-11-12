package models

type PerfData struct {
	MaxPacketDelay    *int   `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int   `json:"avgPacketLossRate,omitempty"`
	AvgTrafficRate    string `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string `json:"maxTrafficRate,omitempty"`
	AvePacketDelay    *int   `json:"avePacketDelay,omitempty"`
}
