/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotificationInfo struct {
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
	NotifId     string `json:"notifId"`
	NotifUri    string `json:"notifUri"`
}