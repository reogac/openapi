package models
type UeContextCancelRelocateData struct {
	 Supi	string	`json:"supi,omitempty"`
	 RelocationCancelRequest	RefToBinaryData	`json:"relocationCancelRequest"`
}
