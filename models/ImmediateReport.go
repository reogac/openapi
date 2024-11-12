package models

type ImmediateReport struct {
	SharedData           []SharedData          `json:"SharedData,omitempty"`
	SubscriptionDataSets *SubscriptionDataSets `json:"SubscriptionDataSets,omitempty"`
}
