/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeEapSession struct {
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	Nonce2            string          `json:"nonce2,omitempty"`
	FiveGPrukId       string          `json:"5gPrukId,omitempty"`
	EapPayload        string          `json:"eapPayload"`
	KnrProSe          string          `json:"knrProSe,omitempty"`
	Links             map[string]Link `json:"_links,omitempty"`
	AuthResult        AuthResult      `json:"authResult,omitempty"`
}