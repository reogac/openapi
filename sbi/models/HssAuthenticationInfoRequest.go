/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationInfoRequest struct {
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	HssAuthType           HssAuthType            `json:"hssAuthType"`
	NumOfRequestedVectors int                    `json:"numOfRequestedVectors"`
	RequestingNodeType    NodeType               `json:"requestingNodeType,omitempty"`
	ServingNetworkId      *PlmnId                `json:"servingNetworkId,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	AnId                  AccessNetworkId        `json:"anId,omitempty"`
}
