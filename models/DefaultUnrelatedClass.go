package models

type DefaultUnrelatedClass struct {
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd               CodeWordInd               `json:"codeWordInd,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	CodeWordList              []string                  `json:"codeWordList,omitempty"`
}