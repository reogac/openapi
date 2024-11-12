package models

type EthFlowDescription struct {
	FDesc          string   `json:"fDesc,omitempty"`
	FDir           string   `json:"fDir,omitempty"`
	SourceMacAddr  string   `json:"sourceMacAddr,omitempty"`
	VlanTags       []string `json:"vlanTags,omitempty"`
	SrcMacAddrEnd  string   `json:"srcMacAddrEnd,omitempty"`
	DestMacAddrEnd string   `json:"destMacAddrEnd,omitempty"`
	DestMacAddr    string   `json:"destMacAddr,omitempty"`
	EthType        string   `json:"ethType"`
}
