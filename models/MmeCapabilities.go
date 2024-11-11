package models
type MmeCapabilities struct {
	 EthernetSupported	*bool	`json:"ethernetSupported,omitempty"`
	 NonIpSupported	*bool	`json:"nonIpSupported,omitempty"`
}
