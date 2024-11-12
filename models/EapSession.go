package models

type EapSession struct {
	SupportedFeatures string                 `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	Msk               string                 `json:"msk,omitempty"`
	EapPayload        string                 `json:"eapPayload"`
	KSeaf             string                 `json:"kSeaf,omitempty"`
	Links             map[string]Link        `json:"_links,omitempty"`
	AuthResult        AuthResult             `json:"authResult,omitempty"`
	Supi              string                 `json:"supi,omitempty"`
}
