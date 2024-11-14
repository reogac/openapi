/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	KSeaf             string                 `json:"kSeaf,omitempty"`
	Links             map[string]Link        `json:"_links,omitempty"`
	AuthResult        AuthResult             `json:"authResult,omitempty"`
	Supi              string                 `json:"supi,omitempty"`
	SupportedFeatures string                 `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	Msk               string                 `json:"msk,omitempty"`
	EapPayload        string                 `json:"eapPayload"`
}
