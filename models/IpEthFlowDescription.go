package models

type IpEthFlowDescription struct {
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty"`
	IpTrafficFilter  string              `json:"ipTrafficFilter,omitempty"`
}
