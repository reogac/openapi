package models
type QosFlowNotifyItem struct {
	 Qfi	int	`json:"qfi"`
	 NotificationCause	string	`json:"notificationCause"`
	 CurrentQosProfileIndex	*int	`json:"currentQosProfileIndex,omitempty"`
	 NullQoSProfileIndex	*bool	`json:"nullQoSProfileIndex,omitempty"`
}
