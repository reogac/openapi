/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdpSubscriptionData struct {
	UpdpNotifySubscriptionId string `json:"updpNotifySubscriptionId"`
	UpdpNotifyCallbackUri    string `json:"updpNotifyCallbackUri"`
	SupportedFeatures        string `json:"supportedFeatures,omitempty"`
	UpdpCallbackBinding      string `json:"updpCallbackBinding,omitempty"`
}
