package models
type QosFlowNotifyItem struct {
	 Qfi	int	`json:"qfi"`
	 NotificationCause	NotificationCause	`json:"notificationCause"`
	 CurrentQosProfileIndex	*int	`json:"currentQosProfileIndex,omitempty"`
	 NullQoSProfileIndex	*bool	`json:"nullQoSProfileIndex,omitempty"`
}
