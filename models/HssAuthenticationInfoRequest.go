package models

type HssAuthenticationInfoRequest struct {
	AnId                  AccessNetworkId        `json:"anId,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	HssAuthType           HssAuthType            `json:"hssAuthType"`
	NumOfRequestedVectors int                    `json:"numOfRequestedVectors"`
	RequestingNodeType    NodeType               `json:"requestingNodeType,omitempty"`
	ServingNetworkId      *PlmnId                `json:"servingNetworkId,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
}
