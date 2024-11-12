package models

type N3gaLocation struct {
	HfcNodeId      *HfcNodeId `json:"hfcNodeId,omitempty"`
	N3IwfId        string     `json:"n3IwfId,omitempty"`
	UeIpv4Addr     string     `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr     string     `json:"ueIpv6Addr,omitempty"`
	PortNumber     *int       `json:"portNumber,omitempty"`
	Gli            string     `json:"gli,omitempty"`
	W5gbanLineType string     `json:"w5gbanLineType,omitempty"`
	Gci            string     `json:"gci,omitempty"`
	N3gppTai       *Tai       `json:"n3gppTai,omitempty"`
	Protocol       string     `json:"protocol,omitempty"`
	TnapId         *TnapId    `json:"tnapId,omitempty"`
	TwapId         *TwapId    `json:"twapId,omitempty"`
}
