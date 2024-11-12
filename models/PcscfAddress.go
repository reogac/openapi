package models

type PcscfAddress struct {
	Fqdn      string   `json:"fqdn,omitempty"`
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
}
