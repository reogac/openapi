package models
type StatusNotification struct {
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 TargetDnaiInfo	*TargetDnaiInfo	`json:"targetDnaiInfo,omitempty"`
	 OldPduSessionRef	string	`json:"oldPduSessionRef,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 StatusInfo	StatusInfo	`json:"statusInfo"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 IntraPlmnApiRoot	string	`json:"intraPlmnApiRoot,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
}
