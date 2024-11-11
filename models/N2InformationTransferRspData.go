package models
type N2InformationTransferRspData struct {
	 Result	string	`json:"result"`
	 PwsRspData	*PWSResponseData	`json:"pwsRspData,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
