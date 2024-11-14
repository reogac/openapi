package models
type SmContextStatusNotification struct {
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 NotifyCorrelationIdsForddnFailure	[]string	`json:"notifyCorrelationIdsForddnFailure,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 StatusInfo	StatusInfo	`json:"statusInfo"`
	 DdnFailureStatus	*bool	`json:"ddnFailureStatus,omitempty"`
	 OldSmContextRef	string	`json:"oldSmContextRef,omitempty"`
	 AltAnchorSmfId	string	`json:"altAnchorSmfId,omitempty"`
	 TargetDnaiInfo	*TargetDnaiInfo	`json:"targetDnaiInfo,omitempty"`
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 NewSmfSetId	string	`json:"newSmfSetId,omitempty"`
	 OldSmfId	string	`json:"oldSmfId,omitempty"`
	 AltAnchorSmfUri	string	`json:"altAnchorSmfUri,omitempty"`
	 NewIntermediateSmfId	string	`json:"newIntermediateSmfId,omitempty"`
	 OldPduSessionRef	string	`json:"oldPduSessionRef,omitempty"`
}
