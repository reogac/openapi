package models
type QosMonitoringReport struct {
	 RefPccRuleIds	[]string	`json:"refPccRuleIds"`
	 UlDelays	[]int	`json:"ulDelays,omitempty"`
	 DlDelays	[]int	`json:"dlDelays,omitempty"`
	 RtDelays	[]int	`json:"rtDelays,omitempty"`
	 Pdmf	*bool	`json:"pdmf,omitempty"`
}
