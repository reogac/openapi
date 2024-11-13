package models

type SmsManagementSubscriptionData struct {
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
}
