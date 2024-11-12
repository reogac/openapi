package models

type ProSeEapSession struct {
	KnrProSe          string          `json:"knrProSe,omitempty"`
	_links            map[string]Link `json:"_links,omitempty"`
	AuthResult        AuthResult      `json:"authResult,omitempty"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	Nonce2            string          `json:"nonce2,omitempty"`
	FiveGPrukId       string          `json:"5gPrukId,omitempty"`
	EapPayload        string          `json:"eapPayload"`
}
