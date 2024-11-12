package models

type NwdafRegistrationModification struct {
	NwdafSetId        string   `json:"nwdafSetId,omitempty"`
	AnalyticsIds      []string `json:"analyticsIds,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	NwdafInstanceId   string   `json:"nwdafInstanceId"`
}
