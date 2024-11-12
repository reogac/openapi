package models

type ProSeAuthenticationCtx struct {
	ProSeAuthData     string          `json:"proSeAuthData"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	AuthType          string          `json:"authType"`
	_links            map[string]Link `json:"_links"`
}
