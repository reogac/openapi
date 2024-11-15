/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreatedData struct {
	UeContext            UeContext         `json:"ueContext"`
	TargetToSourceData   N2InfoContent     `json:"targetToSourceData"`
	PduSessionList       []N2SmInformation `json:"pduSessionList"`
	FailedSessionList    []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures    string            `json:"supportedFeatures,omitempty"`
	PcfReselectedInd     *bool             `json:"pcfReselectedInd,omitempty"`
	AnalyticsNotUsedList []string          `json:"analyticsNotUsedList,omitempty"`
}
