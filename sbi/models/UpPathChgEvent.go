/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpPathChgEvent struct {
	NotificationUri string         `json:"notificationUri"`
	NotifCorreId    string         `json:"notifCorreId"`
	DnaiChgType     DnaiChangeType `json:"dnaiChgType"`
	AfAckInd        *bool          `json:"afAckInd,omitempty"`
}