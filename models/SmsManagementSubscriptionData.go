package models

type SmsManagementSubscriptionData struct {
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
}
