package models

type N1N2MessageTransferRspData struct {
	SupportedFeatures string                   `json:"supportedFeatures,omitempty"`
	Cause             N1N2MessageTransferCause `json:"cause"`
}
