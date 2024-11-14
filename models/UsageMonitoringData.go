package models
type UsageMonitoringData struct {
	 UmId	string	`json:"umId"`
	 VolumeThresholdDownlink	*int64	`json:"volumeThresholdDownlink,omitempty"`
	 MonitoringTime	string	`json:"monitoringTime,omitempty"`
	 NextVolThresholdDownlink	*int64	`json:"nextVolThresholdDownlink,omitempty"`
	 NextTimeThreshold	*int	`json:"nextTimeThreshold,omitempty"`
	 InactivityTime	*int	`json:"inactivityTime,omitempty"`
	 ExUsagePccRuleIds	[]string	`json:"exUsagePccRuleIds,omitempty"`
	 VolumeThreshold	*int64	`json:"volumeThreshold,omitempty"`
	 VolumeThresholdUplink	*int64	`json:"volumeThresholdUplink,omitempty"`
	 TimeThreshold	*int	`json:"timeThreshold,omitempty"`
	 NextVolThreshold	*int64	`json:"nextVolThreshold,omitempty"`
	 NextVolThresholdUplink	*int64	`json:"nextVolThresholdUplink,omitempty"`
}
