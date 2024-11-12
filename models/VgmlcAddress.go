package models

type VgmlcAddress struct {
	VgmlcAddressIpv4 string `json:"vgmlcAddressIpv4,omitempty"`
	VgmlcAddressIpv6 string `json:"vgmlcAddressIpv6,omitempty"`
	VgmlcFqdn        string `json:"vgmlcFqdn,omitempty"`
}
