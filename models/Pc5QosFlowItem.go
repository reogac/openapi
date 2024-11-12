package models

type Pc5QosFlowItem struct {
	Pc5FlowBitRates *Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`
	Range           *int             `json:"range,omitempty"`
	Pqi             int              `json:"pqi"`
}
