package models

type ReportItem struct {
	Path   string `json:"path"`
	Reason string `json:"reason,omitempty"`
}
