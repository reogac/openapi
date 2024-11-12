package models

type WlanPerTsPerformanceInfo struct {
	NumberOfUes *int                `json:"numberOfUes,omitempty"`
	Confidence  *int                `json:"confidence,omitempty"`
	TsStart     string              `json:"tsStart"`
	TsDuration  int                 `json:"tsDuration"`
	Rssi        *int                `json:"rssi,omitempty"`
	Rtt         *int                `json:"rtt,omitempty"`
	TrafficInfo *TrafficInformation `json:"trafficInfo,omitempty"`
}
