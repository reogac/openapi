package models

type ConfirmationDataResponse struct {
	Kseaf      string                 `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	AuthResult AuthResult             `json:"authResult"`
	Supi       string                 `json:"supi,omitempty"`
}
