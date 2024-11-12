package models

type RoutingInfoSmRequest struct {
	IpSmGwInd         *bool  `json:"ipSmGwInd,omitempty"`
	CorrelationId     string `json:"correlationId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
