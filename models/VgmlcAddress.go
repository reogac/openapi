package models

type VgmlcAddress struct {
	VgmlcFqdn        string `json:"vgmlcFqdn,omitempty"`
	VgmlcAddressIpv4 string `json:"vgmlcAddressIpv4,omitempty"`
	VgmlcAddressIpv6 string `json:"vgmlcAddressIpv6,omitempty"`
}
