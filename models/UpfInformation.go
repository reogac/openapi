package models
type UpfInformation struct {
	 UpfId	string	`json:"upfId,omitempty"`
	 UpfAddr	*AddrFqdn	`json:"upfAddr,omitempty"`
}
