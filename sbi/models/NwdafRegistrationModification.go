/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NwdafRegistrationModification struct {
	NwdafSetId        string   `json:"nwdafSetId,omitempty"`
	AnalyticsIds      []string `json:"analyticsIds,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	NwdafInstanceId   string   `json:"nwdafInstanceId"`
}
