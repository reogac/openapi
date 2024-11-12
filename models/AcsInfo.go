package models

type AcsInfo struct {
	AcsIpv6Addr string `json:"acsIpv6Addr,omitempty"`
	AcsUrl      string `json:"acsUrl,omitempty"`
	AcsIpv4Addr string `json:"acsIpv4Addr,omitempty"`
}
