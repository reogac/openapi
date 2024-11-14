package models
type NotificationInfo struct {
	 NotifId	string	`json:"notifId"`
	 NotifUri	string	`json:"notifUri"`
	 UpBufferInd	*bool	`json:"upBufferInd,omitempty"`
}
