package models
type ServerAddressingInfo struct {
	 FqdnList	[]string	`json:"fqdnList,omitempty"`
	 Ipv4Addresses	[]string	`json:"ipv4Addresses,omitempty"`
	 Ipv6Addresses	[]string	`json:"ipv6Addresses,omitempty"`
}
