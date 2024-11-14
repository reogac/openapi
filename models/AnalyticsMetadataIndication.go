/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AnalyticsMetadataIndication struct {
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	AggrNwdafIds  []string       `json:"aggrNwdafIds,omitempty"`
}
