package models
type SecondaryRatUsageInfo struct {
	 SecondaryRatType	string	`json:"secondaryRatType"`
	 QosFlowsUsageData	[]QosFlowUsageReport	`json:"qosFlowsUsageData,omitempty"`
	 PduSessionUsageData	[]VolumeTimedReport	`json:"pduSessionUsageData,omitempty"`
}
