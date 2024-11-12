package models

type RgAuthCtx struct {
	AuthInd    *bool      `json:"authInd,omitempty"`
	AuthResult AuthResult `json:"authResult"`
	Supi       string     `json:"supi,omitempty"`
}
