package models

type FlowInformation struct {
	FlowDescription    string              `json:"flowDescription,omitempty"`
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty"`
	PackFiltId         string              `json:"packFiltId,omitempty"`
	PacketFilterUsage  *bool               `json:"packetFilterUsage,omitempty"`
	TosTrafficClass    string              `json:"tosTrafficClass,omitempty"`
	Spi                string              `json:"spi,omitempty"`
	FlowLabel          string              `json:"flowLabel,omitempty"`
	FlowDirection      FlowDirection       `json:"flowDirection,omitempty"`
}
