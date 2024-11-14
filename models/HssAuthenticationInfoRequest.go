/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationInfoRequest struct {
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	AnId                  AccessNetworkId        `json:"anId,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	HssAuthType           HssAuthType            `json:"hssAuthType"`
	NumOfRequestedVectors int                    `json:"numOfRequestedVectors"`
	RequestingNodeType    NodeType               `json:"requestingNodeType,omitempty"`
	ServingNetworkId      *PlmnId                `json:"servingNetworkId,omitempty"`
}
