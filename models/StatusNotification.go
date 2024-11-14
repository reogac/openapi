package models
type StatusNotification struct {
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 StatusInfo	StatusInfo	`json:"statusInfo"`
	 OldPduSessionRef	string	`json:"oldPduSessionRef,omitempty"`
	 IntraPlmnApiRoot	string	`json:"intraPlmnApiRoot,omitempty"`
	 TargetDnaiInfo	*TargetDnaiInfo	`json:"targetDnaiInfo,omitempty"`
}
