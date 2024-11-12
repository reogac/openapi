package models

type UsageMonitoringData struct {
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	UmId                     string   `json:"umId"`
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
}
