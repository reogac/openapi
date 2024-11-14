/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonitoringData struct {
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	UmId                     string   `json:"umId"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
}
