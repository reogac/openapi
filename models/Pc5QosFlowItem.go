package models

type Pc5QosFlowItem struct {
	Pqi             int              `json:"pqi"`
	Pc5FlowBitRates *Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`
	Range           *int             `json:"range,omitempty"`
}
