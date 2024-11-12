package models

type IpMulticastAddressInfo struct {
	Ipv4MulAddr string `json:"ipv4MulAddr,omitempty"`
	SrcIpv6Addr string `json:"srcIpv6Addr,omitempty"`
	Ipv6MulAddr string `json:"ipv6MulAddr,omitempty"`
	SrcIpv4Addr string `json:"srcIpv4Addr,omitempty"`
}
