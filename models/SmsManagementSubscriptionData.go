package models

type SmsManagementSubscriptionData struct {
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
}
