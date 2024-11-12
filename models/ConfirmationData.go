package models

type ConfirmationData struct {
	ResStar           string `json:"resStar"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
