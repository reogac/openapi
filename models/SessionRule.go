package models

type SessionRule struct {
	SessRuleId   string                `json:"sessRuleId"`
	RefUmData    string                `json:"refUmData,omitempty"`
	RefUmN3gData string                `json:"refUmN3gData,omitempty"`
	RefCondData  string                `json:"refCondData,omitempty"`
	AuthSessAmbr *Ambr                 `json:"authSessAmbr,omitempty"`
	AuthDefQos   *AuthorizedDefaultQos `json:"authDefQos,omitempty"`
}
