package models
type TngfInfo struct {
	 Ipv4EndpointAddresses	[]string	`json:"ipv4EndpointAddresses,omitempty"`
	 Ipv6EndpointAddresses	[]string	`json:"ipv6EndpointAddresses,omitempty"`
	 EndpointFqdn	string	`json:"endpointFqdn,omitempty"`
}
