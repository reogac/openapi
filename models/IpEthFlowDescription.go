package models

type IpEthFlowDescription struct {
	IpTrafficFilter  string              `json:"ipTrafficFilter,omitempty"`
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}
