/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NwdafRegistration struct {
	NwdafInstanceId   string       `json:"nwdafInstanceId"`
	AnalyticsIds      []string     `json:"analyticsIds"`
	NwdafSetId        string       `json:"nwdafSetId,omitempty"`
	RegistrationTime  string       `json:"registrationTime,omitempty"`
	ContextInfo       *ContextInfo `json:"contextInfo,omitempty"`
	SupportedFeatures string       `json:"supportedFeatures,omitempty"`
	ResetIds          []string     `json:"resetIds,omitempty"`
}
