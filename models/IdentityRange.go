package models

type IdentityRange struct {
	End     string `json:"end,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Start   string `json:"start,omitempty"`
}
