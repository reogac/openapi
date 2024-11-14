/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PrevSubInfo struct {
	SubscriptionId string                         `json:"subscriptionId"`
	NfAnaEvents    []string                       `json:"nfAnaEvents,omitempty"`
	UeAnaEvents    []UeAnalyticsContextDescriptor `json:"ueAnaEvents,omitempty"`
	ProducerId     string                         `json:"producerId,omitempty"`
	ProducerSetId  string                         `json:"producerSetId,omitempty"`
}
