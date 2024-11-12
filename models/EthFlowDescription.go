package models

type EthFlowDescription struct {
	SourceMacAddr  string        `json:"sourceMacAddr,omitempty"`
	VlanTags       []string      `json:"vlanTags,omitempty"`
	SrcMacAddrEnd  string        `json:"srcMacAddrEnd,omitempty"`
	DestMacAddrEnd string        `json:"destMacAddrEnd,omitempty"`
	DestMacAddr    string        `json:"destMacAddr,omitempty"`
	EthType        string        `json:"ethType"`
	FDesc          string        `json:"fDesc,omitempty"`
	FDir           FlowDirection `json:"fDir,omitempty"`
}
