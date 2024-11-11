package models
type TwifInfo struct {
	 EndpointFqdn	string	`json:"endpointFqdn,omitempty"`
	 Ipv4EndpointAddresses	[]string	`json:"ipv4EndpointAddresses,omitempty"`
	 Ipv6EndpointAddresses	[]string	`json:"ipv6EndpointAddresses,omitempty"`
}
