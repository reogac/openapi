package models

type IpAddress struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
}
