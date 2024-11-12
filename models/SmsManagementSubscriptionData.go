package models

type SmsManagementSubscriptionData struct {
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
}
