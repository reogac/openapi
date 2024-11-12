package models

type DeregistrationInfo struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	Supi              string `json:"supi"`
}
