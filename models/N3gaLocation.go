package models

type N3gaLocation struct {
	N3IwfId        string     `json:"n3IwfId,omitempty"`
	UeIpv4Addr     string     `json:"ueIpv4Addr,omitempty"`
	TwapId         *TwapId    `json:"twapId,omitempty"`
	W5gbanLineType string     `json:"w5gbanLineType,omitempty"`
	N3gppTai       *Tai       `json:"n3gppTai,omitempty"`
	UeIpv6Addr     string     `json:"ueIpv6Addr,omitempty"`
	PortNumber     *int       `json:"portNumber,omitempty"`
	TnapId         *TnapId    `json:"tnapId,omitempty"`
	Protocol       string     `json:"protocol,omitempty"`
	HfcNodeId      *HfcNodeId `json:"hfcNodeId,omitempty"`
	Gli            string     `json:"gli,omitempty"`
	Gci            string     `json:"gci,omitempty"`
}
