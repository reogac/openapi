package models

type ProSeAuthenticationInfoResult struct {
	AuthType                   AuthType        `json:"authType"`
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
	Supi                       string          `json:"supi,omitempty"`
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
}
