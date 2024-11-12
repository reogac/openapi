package models

type ProSeAuthenticationCtx struct {
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	AuthType          AuthType        `json:"authType"`
	Links             map[string]Link `json:"_links"`
	ProSeAuthData     string          `json:"proSeAuthData"`
}
