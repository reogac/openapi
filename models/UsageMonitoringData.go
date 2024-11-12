package models

type UsageMonitoringData struct {
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	UmId                     string   `json:"umId"`
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
}
