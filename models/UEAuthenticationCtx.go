package models

type UEAuthenticationCtx struct {
	AuthType           AuthType        `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
	Links              map[string]Link `json:"_links"`
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
}
