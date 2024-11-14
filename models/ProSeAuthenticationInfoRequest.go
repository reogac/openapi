package models
type ProSeAuthenticationInfoRequest struct {
	 ServingNetworkName	string	`json:"servingNetworkName"`
	 RelayServiceCode	int	`json:"relayServiceCode"`
	 ResynchronizationInfo	*ResynchronizationInfo	`json:"resynchronizationInfo,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
