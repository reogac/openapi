package models

type UsageMonitoringData struct {
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	UmId                     string   `json:"umId"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
}
