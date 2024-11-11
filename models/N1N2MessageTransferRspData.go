package models

type N1N2MessageTransferRspData struct {
	Cause             string `json:"cause"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
