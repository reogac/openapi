/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfo struct {
	SupiOrSuci         string `json:"supiOrSuci,omitempty"`
	FiveGPrukId        string `json:"5gPrukId,omitempty"`
	RelayServiceCode   int    `json:"relayServiceCode"`
	Nonce1             string `json:"nonce1"`
	ServingNetworkName string `json:"servingNetworkName"`
	SupportedFeatures  string `json:"supportedFeatures,omitempty"`
}
