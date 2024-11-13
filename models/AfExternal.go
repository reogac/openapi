package models

type AfExternal struct {
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	AfId                      string                    `json:"afId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
}
