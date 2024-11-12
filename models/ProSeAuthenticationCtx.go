package models

type ProSeAuthenticationCtx struct {
	AuthType          AuthType        `json:"authType"`
	_links            map[string]Link `json:"_links"`
	ProSeAuthData     string          `json:"proSeAuthData"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
}
