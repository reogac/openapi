package models
type QosFlowNotifyItem struct {
	 CurrentQosProfileIndex	*int	`json:"currentQosProfileIndex,omitempty"`
	 NullQoSProfileIndex	*bool	`json:"nullQoSProfileIndex,omitempty"`
	 Qfi	int	`json:"qfi"`
	 NotificationCause	NotificationCause	`json:"notificationCause"`
}
