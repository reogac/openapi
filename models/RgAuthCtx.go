package models

type RgAuthCtx struct {
	AuthResult AuthResult `json:"authResult"`
	Supi       string     `json:"supi,omitempty"`
	AuthInd    *bool      `json:"authInd,omitempty"`
}
