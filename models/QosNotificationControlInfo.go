package models

type QosNotificationControlInfo struct {
	RefPccRuleIds []string     `json:"refPccRuleIds"`
	NotifType     QosNotifType `json:"notifType"`
	ContVer       *int         `json:"contVer,omitempty"`
	AltQosParamId string       `json:"altQosParamId,omitempty"`
}
