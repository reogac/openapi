package models

type N2InformationTransferRspData struct {
	Result            N2InformationTransferResult `json:"result"`
	PwsRspData        *PWSResponseData            `json:"pwsRspData,omitempty"`
	SupportedFeatures string                      `json:"supportedFeatures,omitempty"`
}
