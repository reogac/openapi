package models

type SmallDataRateStatus struct {
	RemainPacketsDl   *int   `json:"remainPacketsDl,omitempty"`
	ValidityTime      string `json:"validityTime,omitempty"`
	RemainExReportsUl *int   `json:"remainExReportsUl,omitempty"`
	RemainExReportsDl *int   `json:"remainExReportsDl,omitempty"`
	RemainPacketsUl   *int   `json:"remainPacketsUl,omitempty"`
}
