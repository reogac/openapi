package models
type SmContextStatusNotification struct {
	 NewIntermediateSmfId	string	`json:"newIntermediateSmfId,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 OldSmfId	string	`json:"oldSmfId,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 DdnFailureStatus	*bool	`json:"ddnFailureStatus,omitempty"`
	 AltAnchorSmfId	string	`json:"altAnchorSmfId,omitempty"`
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 NewSmfSetId	string	`json:"newSmfSetId,omitempty"`
	 AltAnchorSmfUri	string	`json:"altAnchorSmfUri,omitempty"`
	 TargetDnaiInfo	*TargetDnaiInfo	`json:"targetDnaiInfo,omitempty"`
	 StatusInfo	StatusInfo	`json:"statusInfo"`
	 OldSmContextRef	string	`json:"oldSmContextRef,omitempty"`
	 OldPduSessionRef	string	`json:"oldPduSessionRef,omitempty"`
	 NotifyCorrelationIdsForddnFailure	[]string	`json:"notifyCorrelationIdsForddnFailure,omitempty"`
}
