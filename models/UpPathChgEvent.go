package models

type UpPathChgEvent struct {
	NotificationUri string         `json:"notificationUri"`
	NotifCorreId    string         `json:"notifCorreId"`
	DnaiChgType     DnaiChangeType `json:"dnaiChgType"`
	AfAckInd        *bool          `json:"afAckInd,omitempty"`
}
