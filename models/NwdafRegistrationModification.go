package models
type NwdafRegistrationModification struct {
	 NwdafInstanceId	string	`json:"nwdafInstanceId"`
	 NwdafSetId	string	`json:"nwdafSetId,omitempty"`
	 AnalyticsIds	[]string	`json:"analyticsIds,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
