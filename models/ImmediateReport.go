package models
type ImmediateReport struct {
	 SubscriptionDataSets	*SubscriptionDataSets	`json:"SubscriptionDataSets,omitempty"`
	 SharedData	[]SharedData	`json:"SharedData,omitempty"`
}
