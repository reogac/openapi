package models
type TngfInfo struct {
	 Ipv6EndpointAddresses	[]string	`json:"ipv6EndpointAddresses,omitempty"`
	 EndpointFqdn	string	`json:"endpointFqdn,omitempty"`
	 Ipv4EndpointAddresses	[]string	`json:"ipv4EndpointAddresses,omitempty"`
}
