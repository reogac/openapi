package models

type DnPerformanceReq struct {
	Order             MatchingDirection       `json:"order,omitempty"`
	ReportThresholds  []ThresholdLevel        `json:"reportThresholds,omitempty"`
	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty"`
}
