/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceTypeUnrelatedClass struct {
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	CodeWordList              []string                  `json:"codeWordList,omitempty"`
	ServiceType               int                       `json:"serviceType"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd               CodeWordInd               `json:"codeWordInd,omitempty"`
}
