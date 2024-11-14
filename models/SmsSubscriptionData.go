/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsSubscriptionData struct {
	SupportedFeatures   string `json:"supportedFeatures,omitempty"`
	SmsSubscribed       *bool  `json:"smsSubscribed,omitempty"`
	SharedSmsSubsDataId string `json:"sharedSmsSubsDataId,omitempty"`
}
