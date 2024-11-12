package models

type DeregistrationInfo struct {
	Supi              string `json:"supi"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
