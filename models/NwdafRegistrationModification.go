/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NwdafRegistrationModification struct {
	NwdafInstanceId   string   `json:"nwdafInstanceId"`
	NwdafSetId        string   `json:"nwdafSetId,omitempty"`
	AnalyticsIds      []string `json:"analyticsIds,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
}
