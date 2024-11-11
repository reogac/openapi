package models

type RetainabilityThreshold struct {
	RelFlowNum   *int   `json:"relFlowNum,omitempty"`
	RelTimeUnit  string `json:"relTimeUnit,omitempty"`
	RelFlowRatio *int   `json:"relFlowRatio,omitempty"`
}
