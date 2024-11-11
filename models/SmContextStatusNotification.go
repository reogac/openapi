type SmContextStatusNotification struct {
	 OldSmfId	string	`json:"oldSmfId,omitempty"`
	 AltAnchorSmfId	string	`json:"altAnchorSmfId,omitempty"`
	 StatusInfo	StatusInfo	`json:"statusInfo"`
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 DdnFailureStatus	*bool	`json:"ddnFailureStatus,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 NotifyCorrelationIdsForddnFailure	[]string	`json:"notifyCorrelationIdsForddnFailure,omitempty"`
	 NewIntermediateSmfId	string	`json:"newIntermediateSmfId,omitempty"`
	 NewSmfSetId	string	`json:"newSmfSetId,omitempty"`
	 OldSmContextRef	string	`json:"oldSmContextRef,omitempty"`
	 AltAnchorSmfUri	string	`json:"altAnchorSmfUri,omitempty"`
}
