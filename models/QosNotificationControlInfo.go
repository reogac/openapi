/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosNotificationControlInfo struct {
	NotifType     QosNotifType `json:"notifType"`
	ContVer       *int         `json:"contVer,omitempty"`
	AltQosParamId string       `json:"altQosParamId,omitempty"`
	RefPccRuleIds []string     `json:"refPccRuleIds"`
}
