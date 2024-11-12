package models

type IdTranslationResult struct {
	Supi              string   `json:"supi"`
	Gpsi              string   `json:"gpsi,omitempty"`
	AdditionalSupis   []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis   []string `json:"additionalGpsis,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
}
