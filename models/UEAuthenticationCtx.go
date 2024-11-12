package models

type UEAuthenticationCtx struct {
	AuthType           string          `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
	_links             map[string]Link `json:"_links"`
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
}
