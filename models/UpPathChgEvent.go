package models

type UpPathChgEvent struct {
	NotifCorreId    string         `json:"notifCorreId"`
	DnaiChgType     DnaiChangeType `json:"dnaiChgType"`
	AfAckInd        *bool          `json:"afAckInd,omitempty"`
	NotificationUri string         `json:"notificationUri"`
}
