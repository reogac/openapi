package models
type SmsManagementSubscriptionData struct {
	 MtSmsSubscribed	*bool	`json:"mtSmsSubscribed,omitempty"`
	 MtSmsBarringRoaming	*bool	`json:"mtSmsBarringRoaming,omitempty"`
	 MoSmsBarringAll	*bool	`json:"moSmsBarringAll,omitempty"`
	 SharedSmsMngDataIds	[]string	`json:"sharedSmsMngDataIds,omitempty"`
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 MtSmsBarringAll	*bool	`json:"mtSmsBarringAll,omitempty"`
	 MoSmsSubscribed	*bool	`json:"moSmsSubscribed,omitempty"`
	 MoSmsBarringRoaming	*bool	`json:"moSmsBarringRoaming,omitempty"`
}
