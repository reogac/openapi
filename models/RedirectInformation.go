package models

type RedirectInformation struct {
	RedirectEnabled       *bool               `json:"redirectEnabled,omitempty"`
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType,omitempty"`
	RedirectServerAddress string              `json:"redirectServerAddress,omitempty"`
}
