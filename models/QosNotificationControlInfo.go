package models
type QosNotificationControlInfo struct {
	 NotifType	QosNotifType	`json:"notifType"`
	 ContVer	*int	`json:"contVer,omitempty"`
	 AltQosParamId	string	`json:"altQosParamId,omitempty"`
	 RefPccRuleIds	[]string	`json:"refPccRuleIds"`
}
