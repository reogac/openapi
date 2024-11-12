package models
type ServerAddressingInfo struct {
	 Ipv4Addresses	[]string	`json:"ipv4Addresses,omitempty"`
	 Ipv6Addresses	[]string	`json:"ipv6Addresses,omitempty"`
	 FqdnList	[]string	`json:"fqdnList,omitempty"`
}
