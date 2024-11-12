package models
type AnalyticsSubscription struct {
	 NwdafId	string	`json:"nwdafId,omitempty"`
	 NwdafSetId	string	`json:"nwdafSetId,omitempty"`
	 NwdafSubscriptionList	[]NwdafSubscription	`json:"nwdafSubscriptionList"`
}
