/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotificationInfo struct {
	NotifId     string `json:"notifId"`
	NotifUri    string `json:"notifUri"`
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
}
