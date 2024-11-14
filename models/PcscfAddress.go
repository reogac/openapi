package models
type PcscfAddress struct {
	 Ipv4Addrs	[]string	`json:"ipv4Addrs,omitempty"`
	 Ipv6Addrs	[]string	`json:"ipv6Addrs,omitempty"`
	 Fqdn	string	`json:"fqdn,omitempty"`
}
