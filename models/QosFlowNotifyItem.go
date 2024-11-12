package models

type QosFlowNotifyItem struct {
	NotificationCause      NotificationCause `json:"notificationCause"`
	CurrentQosProfileIndex *int              `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool             `json:"nullQoSProfileIndex,omitempty"`
	Qfi                    int               `json:"qfi"`
}
