package models

type DnPerformanceReq struct {
	DnPerfOrderCriter string           `json:"dnPerfOrderCriter,omitempty"`
	Order             string           `json:"order,omitempty"`
	ReportThresholds  []ThresholdLevel `json:"reportThresholds,omitempty"`
}
