type SecondaryRatUsageReport struct {
	 SecondaryRatType	string	`json:"secondaryRatType"`
	 QosFlowsUsageData	[]QosFlowUsageReport	`json:"qosFlowsUsageData"`
}
