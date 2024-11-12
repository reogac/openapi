package models

type SorUpdateInfo struct {
	VplmnId           PlmnId `json:"vplmnId"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
