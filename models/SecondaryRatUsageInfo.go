package models
type SecondaryRatUsageInfo struct {
	 PduSessionUsageData	[]VolumeTimedReport	`json:"pduSessionUsageData,omitempty"`
	 SecondaryRatType	RatType	`json:"secondaryRatType"`
	 QosFlowsUsageData	[]QosFlowUsageReport	`json:"qosFlowsUsageData,omitempty"`
}
