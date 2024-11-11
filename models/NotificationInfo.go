package models

type NotificationInfo struct {
	NotifUri    string `json:"notifUri"`
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
	NotifId     string `json:"notifId"`
}
