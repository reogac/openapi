package models

type IdentityRange struct {
	Pattern string `json:"pattern,omitempty"`
	Start   string `json:"start,omitempty"`
	End     string `json:"end,omitempty"`
}
