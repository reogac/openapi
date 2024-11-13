package models

type N2InformationTransferRspData struct {
	SupportedFeatures string                      `json:"supportedFeatures,omitempty"`
	Result            N2InformationTransferResult `json:"result"`
	PwsRspData        *PWSResponseData            `json:"pwsRspData,omitempty"`
}
