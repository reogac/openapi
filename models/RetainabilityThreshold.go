package models

type RetainabilityThreshold struct {
	RelFlowRatio *int     `json:"relFlowRatio,omitempty"`
	RelFlowNum   *int     `json:"relFlowNum,omitempty"`
	RelTimeUnit  TimeUnit `json:"relTimeUnit,omitempty"`
}
