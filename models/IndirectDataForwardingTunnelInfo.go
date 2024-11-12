package models

type IndirectDataForwardingTunnelInfo struct {
	AdditionalTnlNb *int   `json:"additionalTnlNb,omitempty"`
	Ipv4Addr        string `json:"ipv4Addr,omitempty"`
	Ipv6Addr        string `json:"ipv6Addr,omitempty"`
	GtpTeid         string `json:"gtpTeid"`
	DrbId           *int   `json:"drbId,omitempty"`
}
