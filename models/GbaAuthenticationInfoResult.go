package models

type GbaAuthenticationInfoResult struct {
	ThreeGAkaAv       *ThreeGAkaAv `json:"3gAkaAv,omitempty"`
	SupportedFeatures string       `json:"supportedFeatures,omitempty"`
}
