package models

type QosNotificationControlInfo struct {
	ContVer       *int         `json:"contVer,omitempty"`
	AltQosParamId string       `json:"altQosParamId,omitempty"`
	RefPccRuleIds []string     `json:"refPccRuleIds"`
	NotifType     QosNotifType `json:"notifType"`
}
