package models

type AddrFqdn struct {
	IpAddr *IpAddr `json:"ipAddr,omitempty"`
	Fqdn   string  `json:"fqdn,omitempty"`
}
