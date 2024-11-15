/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeEapSession struct {
	AuthResult        AuthResult      `json:"authResult,omitempty"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	Nonce2            string          `json:"nonce2,omitempty"`
	FiveGPrukId       string          `json:"5gPrukId,omitempty"`
	EapPayload        string          `json:"eapPayload"`
	KnrProSe          string          `json:"knrProSe,omitempty"`
	Links             map[string]Link `json:"_links,omitempty"`
}
