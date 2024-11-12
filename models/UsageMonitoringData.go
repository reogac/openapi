package models

type UsageMonitoringData struct {
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	UmId                     string   `json:"umId"`
}
