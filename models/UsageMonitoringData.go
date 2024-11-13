package models

type UsageMonitoringData struct {
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	UmId                     string   `json:"umId"`
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
}
