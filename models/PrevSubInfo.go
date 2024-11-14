package models
type PrevSubInfo struct {
	 ProducerId	string	`json:"producerId,omitempty"`
	 ProducerSetId	string	`json:"producerSetId,omitempty"`
	 SubscriptionId	string	`json:"subscriptionId"`
	 NfAnaEvents	[]string	`json:"nfAnaEvents,omitempty"`
	 UeAnaEvents	[]UeAnalyticsContextDescriptor	`json:"ueAnaEvents,omitempty"`
}
