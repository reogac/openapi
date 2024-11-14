package models
type WlanPerTsPerformanceInfo struct {
	 Confidence	*int	`json:"confidence,omitempty"`
	 TsStart	string	`json:"tsStart"`
	 TsDuration	int	`json:"tsDuration"`
	 Rssi	*int	`json:"rssi,omitempty"`
	 Rtt	*int	`json:"rtt,omitempty"`
	 TrafficInfo	*TrafficInformation	`json:"trafficInfo,omitempty"`
	 NumberOfUes	*int	`json:"numberOfUes,omitempty"`
}
