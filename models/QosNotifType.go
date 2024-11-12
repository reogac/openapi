package models

type QosNotifType string

// Define constant values for QosNotifType
const (
	QOSNOTIFTYPE_GUARANTEED     QosNotifType = "GUARANTEED"
	QOSNOTIFTYPE_NOT_GUARANTEED QosNotifType = "NOT_GUARANTEED"
)
