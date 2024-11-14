package models
type RoutingInfoSmRequest struct {
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 IpSmGwInd	*bool	`json:"ipSmGwInd,omitempty"`
	 CorrelationId	string	`json:"correlationId,omitempty"`
}
