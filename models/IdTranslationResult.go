package models

type IdTranslationResult struct {
	AdditionalGpsis   []string `json:"additionalGpsis,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	Supi              string   `json:"supi"`
	Gpsi              string   `json:"gpsi,omitempty"`
	AdditionalSupis   []string `json:"additionalSupis,omitempty"`
}
