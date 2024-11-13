package models

type N3gaLocation struct {
	Gci            string            `json:"gci,omitempty"`
	N3IwfId        string            `json:"n3IwfId,omitempty"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty"`
	Protocol       TransportProtocol `json:"protocol,omitempty"`
	TwapId         *TwapId           `json:"twapId,omitempty"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty"`
	N3gppTai       *Tai              `json:"n3gppTai,omitempty"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty"`
	PortNumber     *int              `json:"portNumber,omitempty"`
	TnapId         *TnapId           `json:"tnapId,omitempty"`
	Gli            string            `json:"gli,omitempty"`
}
