/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:09 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationInfoResult struct {
	RoutingId            string                 `json:"routingId,omitempty"`
	PvsInfo              []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	AuthType             AuthType               `json:"authType"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector  `json:"authenticationVector,omitempty"`
	Supi                 string                 `json:"supi,omitempty"`
	AkmaInd              *bool                  `json:"akmaInd,omitempty"`
	AuthAaa              *bool                  `json:"authAaa,omitempty"`
}