package models

type UEAuthenticationCtx struct {
	Links              map[string]Link `json:"_links"`
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
	AuthType           AuthType        `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
}
