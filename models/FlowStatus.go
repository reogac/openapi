package models

type FlowStatus string

// Define constant values for FlowStatus
const (
	FLOWSTATUS_ENABLED_UPLINK   FlowStatus = "ENABLED-UPLINK"
	FLOWSTATUS_ENABLED_DOWNLINK FlowStatus = "ENABLED-DOWNLINK"
	FLOWSTATUS_ENABLED          FlowStatus = "ENABLED"
	FLOWSTATUS_DISABLED         FlowStatus = "DISABLED"
	FLOWSTATUS_REMOVED          FlowStatus = "REMOVED"
)
