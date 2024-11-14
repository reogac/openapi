package models
type ProseSubscriptionData struct {
	 NrUePc5Ambr	string	`json:"nrUePc5Ambr,omitempty"`
	 ProseAllowedPlmn	[]ProSeAllowedPlmn	`json:"proseAllowedPlmn,omitempty"`
	 ProseServiceAuth	*ProseServiceAuth	`json:"proseServiceAuth,omitempty"`
}
