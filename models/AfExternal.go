package models

type AfExternal struct {
	AfId                      string                    `json:"afId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
}
