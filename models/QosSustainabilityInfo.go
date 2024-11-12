package models

type QosSustainabilityInfo struct {
	QosFlowRetThd *RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`
	RanUeThrouThd string                  `json:"ranUeThrouThd,omitempty"`
	Snssai        *Snssai                 `json:"snssai,omitempty"`
	Confidence    *int                    `json:"confidence,omitempty"`
	AreaInfo      *NetworkAreaInfo        `json:"areaInfo,omitempty"`
	StartTs       string                  `json:"startTs,omitempty"`
	EndTs         string                  `json:"endTs,omitempty"`
}
