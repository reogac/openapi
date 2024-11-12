package models

type AuthenticationInfoResult struct {
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector  `json:"authenticationVector,omitempty"`
	Supi                 string                 `json:"supi,omitempty"`
	AkmaInd              *bool                  `json:"akmaInd,omitempty"`
	AuthAaa              *bool                  `json:"authAaa,omitempty"`
	RoutingId            string                 `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	AuthType             AuthType               `json:"authType"`
}
