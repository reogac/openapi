/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRule struct {
	RefUmN3gData string                `json:"refUmN3gData,omitempty"`
	RefCondData  string                `json:"refCondData,omitempty"`
	AuthSessAmbr *Ambr                 `json:"authSessAmbr,omitempty"`
	AuthDefQos   *AuthorizedDefaultQos `json:"authDefQos,omitempty"`
	SessRuleId   string                `json:"sessRuleId"`
	RefUmData    string                `json:"refUmData,omitempty"`
}
