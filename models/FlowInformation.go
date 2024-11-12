package models

type FlowInformation struct {
	PackFiltId         string              `json:"packFiltId,omitempty"`
	PacketFilterUsage  *bool               `json:"packetFilterUsage,omitempty"`
	TosTrafficClass    string              `json:"tosTrafficClass,omitempty"`
	Spi                string              `json:"spi,omitempty"`
	FlowLabel          string              `json:"flowLabel,omitempty"`
	FlowDirection      FlowDirection       `json:"flowDirection,omitempty"`
	FlowDescription    string              `json:"flowDescription,omitempty"`
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty"`
}
