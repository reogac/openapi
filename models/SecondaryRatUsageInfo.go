package models

type SecondaryRatUsageInfo struct {
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
	SecondaryRatType    RatType              `json:"secondaryRatType"`
}
