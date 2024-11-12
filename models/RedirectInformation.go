package models

type RedirectInformation struct {
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType,omitempty"`
	RedirectServerAddress string              `json:"redirectServerAddress,omitempty"`
	RedirectEnabled       *bool               `json:"redirectEnabled,omitempty"`
}
