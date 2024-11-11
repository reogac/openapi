package models
type ApnRateStatus struct {
	 RemainExReportsUl	*int	`json:"remainExReportsUl,omitempty"`
	 RemainExReportsDl	*int	`json:"remainExReportsDl,omitempty"`
	 RemainPacketsUl	*int	`json:"remainPacketsUl,omitempty"`
	 RemainPacketsDl	*int	`json:"remainPacketsDl,omitempty"`
	 ValidityTime	string	`json:"validityTime,omitempty"`
}
