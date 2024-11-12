package models

type UpPathChgEvent struct {
	AfAckInd        *bool          `json:"afAckInd,omitempty"`
	NotificationUri string         `json:"notificationUri"`
	NotifCorreId    string         `json:"notifCorreId"`
	DnaiChgType     DnaiChangeType `json:"dnaiChgType"`
}
