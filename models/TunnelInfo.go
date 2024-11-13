package models

type TunnelInfo struct {
	GtpTeid  string     `json:"gtpTeid"`
	AnType   AccessType `json:"anType,omitempty"`
	Ipv4Addr string     `json:"ipv4Addr,omitempty"`
	Ipv6Addr string     `json:"ipv6Addr,omitempty"`
}
