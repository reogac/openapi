/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfSelectionSubscriptionData struct {
	HssGroupId            string                `json:"hssGroupId,omitempty"`
	SupportedFeatures     string                `json:"supportedFeatures,omitempty"`
	SubscribedSnssaiInfos map[string]SnssaiInfo `json:"subscribedSnssaiInfos,omitempty"`
	SharedSnssaiInfosId   string                `json:"sharedSnssaiInfosId,omitempty"`
}
