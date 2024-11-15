/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsManagementSubscriptionData struct {
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
}
