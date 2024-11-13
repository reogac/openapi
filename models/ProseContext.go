package models

type ProseContext struct {
	DirectComm      UeAuth      `json:"directComm,omitempty"`
	L2Relay         UeAuth      `json:"l2Relay,omitempty"`
	L3Relay         UeAuth      `json:"l3Relay,omitempty"`
	L2Remote        UeAuth      `json:"l2Remote,omitempty"`
	NrUePc5Ambr     string      `json:"nrUePc5Ambr,omitempty"`
	Pc5QoSPara      *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
	DirectDiscovery UeAuth      `json:"directDiscovery,omitempty"`
}
