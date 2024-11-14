package models
type RouteInformation struct {
	 PortNumber	int	`json:"portNumber"`
	 Ipv4Addr	string	`json:"ipv4Addr,omitempty"`
	 Ipv6Addr	string	`json:"ipv6Addr,omitempty"`
}
