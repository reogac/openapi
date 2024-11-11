package models

type SecondaryRatUsageInfo struct {
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
	SecondaryRatType    string               `json:"secondaryRatType"`
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
}
