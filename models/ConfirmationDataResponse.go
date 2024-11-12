package models
type ConfirmationDataResponse struct {
	 AuthResult	AuthResult	`json:"authResult"`
	 Supi	string	`json:"supi,omitempty"`
	 Kseaf	string	`json:"kseaf,omitempty"`
	 PvsInfo	[]ServerAddressingInfo	`json:"pvsInfo,omitempty"`
}
