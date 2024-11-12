package models

type UpCnxState string

// Define constant values for UpCnxState
const (
	UPCNXSTATE_ACTIVATED   UpCnxState = "ACTIVATED"
	UPCNXSTATE_DEACTIVATED UpCnxState = "DEACTIVATED"
	UPCNXSTATE_ACTIVATING  UpCnxState = "ACTIVATING"
	UPCNXSTATE_SUSPENDED   UpCnxState = "SUSPENDED"
)
