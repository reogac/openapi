package models

type MmeCapabilities struct {
	NonIpSupported    *bool `json:"nonIpSupported,omitempty"`
	EthernetSupported *bool `json:"ethernetSupported,omitempty"`
	UpipSupported     *bool `json:"upipSupported,omitempty"`
}
