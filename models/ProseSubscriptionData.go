package models
type ProseSubscriptionData struct {
	 ProseServiceAuth	*ProseServiceAuth	`json:"proseServiceAuth,omitempty"`
	 NrUePc5Ambr	string	`json:"nrUePc5Ambr,omitempty"`
	 ProseAllowedPlmn	[]ProSeAllowedPlmn	`json:"proseAllowedPlmn,omitempty"`
}
