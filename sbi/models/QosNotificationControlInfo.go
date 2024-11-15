/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosNotificationControlInfo struct {
	RefPccRuleIds []string     `json:"refPccRuleIds"`
	NotifType     QosNotifType `json:"notifType"`
	ContVer       *int         `json:"contVer,omitempty"`
	AltQosParamId string       `json:"altQosParamId,omitempty"`
}
