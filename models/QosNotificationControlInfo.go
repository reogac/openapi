/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosNotificationControlInfo struct {
	ContVer       *int         `json:"contVer,omitempty"`
	AltQosParamId string       `json:"altQosParamId,omitempty"`
	RefPccRuleIds []string     `json:"refPccRuleIds"`
	NotifType     QosNotifType `json:"notifType"`
}
