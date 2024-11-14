package models
type QosNotificationControlInfo struct {
	 AltQosParamId	string	`json:"altQosParamId,omitempty"`
	 RefPccRuleIds	[]string	`json:"refPccRuleIds"`
	 NotifType	QosNotifType	`json:"notifType"`
	 ContVer	*int	`json:"contVer,omitempty"`
}
