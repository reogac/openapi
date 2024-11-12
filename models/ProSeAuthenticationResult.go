package models

type ProSeAuthenticationResult struct {
	KnrProSe          string `json:"knrProSe,omitempty"`
	Nonce2            string `json:"nonce2,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
