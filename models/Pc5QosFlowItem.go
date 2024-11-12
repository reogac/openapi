package models

type Pc5QosFlowItem struct {
	Range           *int             `json:"range,omitempty"`
	Pqi             int              `json:"pqi"`
	Pc5FlowBitRates *Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`
}
