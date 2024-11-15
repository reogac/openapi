/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonUeN2InfoSubscriptionCreateData struct {
	NfId                string             `json:"nfId,omitempty"`
	SupportedFeatures   string             `json:"supportedFeatures,omitempty"`
	GlobalRanNodeList   []GlobalRanNodeId  `json:"globalRanNodeList,omitempty"`
	AnTypeList          []string           `json:"anTypeList,omitempty"`
	N2InformationClass  N2InformationClass `json:"n2InformationClass"`
	N2NotifyCallbackUri string             `json:"n2NotifyCallbackUri"`
}
