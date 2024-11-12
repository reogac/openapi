package models

type SecondaryRatUsageInfo struct {
	SecondaryRatType    RatType              `json:"secondaryRatType"`
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
}
