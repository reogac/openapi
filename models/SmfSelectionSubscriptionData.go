/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfSelectionSubscriptionData struct {
	SharedSnssaiInfosId   string                `json:"sharedSnssaiInfosId,omitempty"`
	HssGroupId            string                `json:"hssGroupId,omitempty"`
	SupportedFeatures     string                `json:"supportedFeatures,omitempty"`
	SubscribedSnssaiInfos map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty"`
}
