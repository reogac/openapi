package models

type UpSecurity struct {
	UpIntegr UpIntegrity       `json:"upIntegr"`
	UpConfid UpConfidentiality `json:"upConfid"`
}
