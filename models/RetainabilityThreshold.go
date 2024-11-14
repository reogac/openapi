package models
type RetainabilityThreshold struct {
	 RelFlowNum	*int	`json:"relFlowNum,omitempty"`
	 RelTimeUnit	TimeUnit	`json:"relTimeUnit,omitempty"`
	 RelFlowRatio	*int	`json:"relFlowRatio,omitempty"`
}
