package models
type SecondaryRatUsageReport struct {
	 SecondaryRatType	RatType	`json:"secondaryRatType"`
	 QosFlowsUsageData	[]QosFlowUsageReport	`json:"qosFlowsUsageData"`
}
