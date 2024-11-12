package models
type UEAuthenticationCtx struct {
	 FiveGAuthData	FiveGAuthData	`json:"5gAuthData"`
	 Links	map[string]Link	`json:"_links"`
	 ServingNetworkName	string	`json:"servingNetworkName,omitempty"`
	 AuthType	AuthType	`json:"authType"`
}
