package models
type AnalyticsSubscription struct {
	 NwdafSetId	string	`json:"nwdafSetId,omitempty"`
	 NwdafSubscriptionList	[]NwdafSubscription	`json:"nwdafSubscriptionList"`
	 NwdafId	string	`json:"nwdafId,omitempty"`
}
