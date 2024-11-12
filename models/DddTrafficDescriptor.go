package models
type DddTrafficDescriptor struct {
	 MacAddr	string	`json:"macAddr,omitempty"`
	 Ipv4Addr	string	`json:"ipv4Addr,omitempty"`
	 Ipv6Addr	string	`json:"ipv6Addr,omitempty"`
	 PortNumber	*int	`json:"portNumber,omitempty"`
}
