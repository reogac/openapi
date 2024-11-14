package models
type MmeCapabilities struct {
	 UpipSupported	*bool	`json:"upipSupported,omitempty"`
	 NonIpSupported	*bool	`json:"nonIpSupported,omitempty"`
	 EthernetSupported	*bool	`json:"ethernetSupported,omitempty"`
}
