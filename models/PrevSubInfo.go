package models

type PrevSubInfo struct {
	NfAnaEvents    []string                       `json:"nfAnaEvents,omitempty"`
	UeAnaEvents    []UeAnalyticsContextDescriptor `json:"ueAnaEvents,omitempty"`
	ProducerId     string                         `json:"producerId,omitempty"`
	ProducerSetId  string                         `json:"producerSetId,omitempty"`
	SubscriptionId string                         `json:"subscriptionId"`
}
