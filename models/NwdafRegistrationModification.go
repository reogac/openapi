package models

type NwdafRegistrationModification struct {
	AnalyticsIds      []string `json:"analyticsIds,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	NwdafInstanceId   string   `json:"nwdafInstanceId"`
	NwdafSetId        string   `json:"nwdafSetId,omitempty"`
}
