package models
type QosSustainabilityInfo struct {
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 AreaInfo	*NetworkAreaInfo	`json:"areaInfo,omitempty"`
	 StartTs	string	`json:"startTs,omitempty"`
	 EndTs	string	`json:"endTs,omitempty"`
	 QosFlowRetThd	*RetainabilityThreshold	`json:"qosFlowRetThd,omitempty"`
	 RanUeThrouThd	string	`json:"ranUeThrouThd,omitempty"`
}
