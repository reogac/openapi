package models

type SorUpdateInfo struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	VplmnId           PlmnId `json:"vplmnId"`
}
