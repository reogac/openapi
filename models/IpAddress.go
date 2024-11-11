package models

type IpAddress struct {
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
}
