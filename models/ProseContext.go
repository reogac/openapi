package models

type ProseContext struct {
	DirectComm      string      `json:"directComm,omitempty"`
	L2Relay         string      `json:"l2Relay,omitempty"`
	L3Relay         string      `json:"l3Relay,omitempty"`
	L2Remote        string      `json:"l2Remote,omitempty"`
	NrUePc5Ambr     string      `json:"nrUePc5Ambr,omitempty"`
	Pc5QoSPara      *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
	DirectDiscovery string      `json:"directDiscovery,omitempty"`
}
