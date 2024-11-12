package models

type ProSeAuthenticationCtx struct {
	ProSeAuthData     string          `json:"proSeAuthData"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	AuthType          AuthType        `json:"authType"`
	Links             map[string]Link `json:"_links"`
}
