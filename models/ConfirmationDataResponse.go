package models

type ConfirmationDataResponse struct {
	Supi       string                 `json:"supi,omitempty"`
	Kseaf      string                 `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo `json:"pvsInfo,omitempty"`
	AuthResult AuthResult             `json:"authResult"`
}
