package models

type ProseSubscriptionData struct {
	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
	ProseServiceAuth *ProseServiceAuth  `json:"proseServiceAuth,omitempty"`
	NrUePc5Ambr      string             `json:"nrUePc5Ambr,omitempty"`
}
