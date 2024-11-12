package models

type TunnelAddress struct {
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	PortNumber int    `json:"portNumber"`
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
}
