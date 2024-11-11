package models

type Pc5QoSPara struct {
	Pc5LinkAmbr    string           `json:"pc5LinkAmbr,omitempty"`
	Pc5QosFlowList []Pc5QosFlowItem `json:"pc5QosFlowList"`
}
