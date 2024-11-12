package models

type NwdafRegistration struct {
	AnalyticsIds      []string     `json:"analyticsIds"`
	NwdafSetId        string       `json:"nwdafSetId,omitempty"`
	RegistrationTime  string       `json:"registrationTime,omitempty"`
	ContextInfo       *ContextInfo `json:"contextInfo,omitempty"`
	SupportedFeatures string       `json:"supportedFeatures,omitempty"`
	ResetIds          []string     `json:"resetIds,omitempty"`
	NwdafInstanceId   string       `json:"nwdafInstanceId"`
}
