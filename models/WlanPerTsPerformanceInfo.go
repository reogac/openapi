package models

type WlanPerTsPerformanceInfo struct {
	Rtt         *int                `json:"rtt,omitempty"`
	TrafficInfo *TrafficInformation `json:"trafficInfo,omitempty"`
	NumberOfUes *int                `json:"numberOfUes,omitempty"`
	Confidence  *int                `json:"confidence,omitempty"`
	TsStart     string              `json:"tsStart"`
	TsDuration  int                 `json:"tsDuration"`
	Rssi        *int                `json:"rssi,omitempty"`
}
