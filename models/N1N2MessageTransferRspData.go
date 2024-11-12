package models
type N1N2MessageTransferRspData struct {
	 Cause	N1N2MessageTransferCause	`json:"cause"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
