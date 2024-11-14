package models
type N3gaLocation struct {
	 UeIpv4Addr	string	`json:"ueIpv4Addr,omitempty"`
	 UeIpv6Addr	string	`json:"ueIpv6Addr,omitempty"`
	 PortNumber	*int	`json:"portNumber,omitempty"`
	 Protocol	TransportProtocol	`json:"protocol,omitempty"`
	 TwapId	*TwapId	`json:"twapId,omitempty"`
	 W5gbanLineType	LineType	`json:"w5gbanLineType,omitempty"`
	 Gci	string	`json:"gci,omitempty"`
	 N3gppTai	*Tai	`json:"n3gppTai,omitempty"`
	 N3IwfId	string	`json:"n3IwfId,omitempty"`
	 TnapId	*TnapId	`json:"tnapId,omitempty"`
	 HfcNodeId	*HfcNodeId	`json:"hfcNodeId,omitempty"`
	 Gli	string	`json:"gli,omitempty"`
}
