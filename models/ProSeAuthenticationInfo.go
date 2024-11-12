package models

type ProSeAuthenticationInfo struct {
	SupportedFeatures  string `json:"supportedFeatures,omitempty"`
	SupiOrSuci         string `json:"supiOrSuci,omitempty"`
	FiveGPrukId        string `json:"5gPrukId,omitempty"`
	RelayServiceCode   int    `json:"relayServiceCode"`
	Nonce1             string `json:"nonce1"`
	ServingNetworkName string `json:"servingNetworkName"`
}
