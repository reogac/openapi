package models

type IdentityRange struct {
	Start   string `json:"start,omitempty"`
	End     string `json:"end,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}
