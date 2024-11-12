package models

type QosMonitoringReport struct {
	UlDelays      []int    `json:"ulDelays,omitempty"`
	DlDelays      []int    `json:"dlDelays,omitempty"`
	RtDelays      []int    `json:"rtDelays,omitempty"`
	Pdmf          *bool    `json:"pdmf,omitempty"`
	RefPccRuleIds []string `json:"refPccRuleIds"`
}
