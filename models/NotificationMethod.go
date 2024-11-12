package models

type NotificationMethod string

// Define constant values for NotificationMethod
const (
	NOTIFICATIONMETHOD_PERIODIC  NotificationMethod = "PERIODIC"
	NOTIFICATIONMETHOD_THRESHOLD NotificationMethod = "THRESHOLD"
)
