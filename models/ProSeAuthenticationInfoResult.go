package models

type ProSeAuthenticationInfoResult struct {
	Supi                       string          `json:"supi,omitempty"`
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
	AuthType                   AuthType        `json:"authType"`
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
}
