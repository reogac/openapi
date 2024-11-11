package models

type MmeCapabilities struct {
	NonIpSupported    *bool `json:"nonIpSupported,omitempty"`
	EthernetSupported *bool `json:"ethernetSupported,omitempty"`
}
