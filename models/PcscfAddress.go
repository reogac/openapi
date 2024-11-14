package models
type PcscfAddress struct {
	 Ipv6Addrs	[]string	`json:"ipv6Addrs,omitempty"`
	 Fqdn	string	`json:"fqdn,omitempty"`
	 Ipv4Addrs	[]string	`json:"ipv4Addrs,omitempty"`
}
