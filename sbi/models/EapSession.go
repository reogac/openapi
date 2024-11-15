/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	Msk               string                 `json:"msk,omitempty"`
	EapPayload        string                 `json:"eapPayload"`
	KSeaf             string                 `json:"kSeaf,omitempty"`
	Links             map[string]Link        `json:"_links,omitempty"`
	AuthResult        AuthResult             `json:"authResult,omitempty"`
	Supi              string                 `json:"supi,omitempty"`
	SupportedFeatures string                 `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo `json:"pvsInfo,omitempty"`
}
