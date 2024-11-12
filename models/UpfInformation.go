package models

type UpfInformation struct {
	UpfAddr *AddrFqdn `json:"upfAddr,omitempty"`
	UpfId   string    `json:"upfId,omitempty"`
}
