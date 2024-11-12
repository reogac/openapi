package models

type AddressList struct {
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
}
