package models

type NotificationCause string

// Define constant values for NotificationCause
const (
	NOTIFICATIONCAUSE_QOS_FULFILLED        NotificationCause = "QOS_FULFILLED"
	NOTIFICATIONCAUSE_QOS_NOT_FULFILLED    NotificationCause = "QOS_NOT_FULFILLED"
	NOTIFICATIONCAUSE_UP_SEC_FULFILLED     NotificationCause = "UP_SEC_FULFILLED"
	NOTIFICATIONCAUSE_UP_SEC_NOT_FULFILLED NotificationCause = "UP_SEC_NOT_FULFILLED"
)