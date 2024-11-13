package models

type CagInfo struct {
	CagOnlyIndicator *bool    `json:"cagOnlyIndicator,omitempty"`
	AllowedCagList   []string `json:"allowedCagList"`
}
