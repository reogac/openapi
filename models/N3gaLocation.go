package models

type N3gaLocation struct {
	Gci            string            `json:"gci,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	Gli            string            `json:"gli,omitempty"`
}
