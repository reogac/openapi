package models

type ProSeAuthenticationInfoResult struct {
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
	Supi                       string          `json:"supi,omitempty"`
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
	AuthType                   AuthType        `json:"authType"`
}
