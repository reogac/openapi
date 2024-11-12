package models
type SecondaryRatUsageReport struct {
	 QosFlowsUsageData	[]QosFlowUsageReport	`json:"qosFlowsUsageData"`
	 SecondaryRatType	string	`json:"secondaryRatType"`
}
