package models

type ResourceStatus string

// Define constant values for ResourceStatus
const (
	RESOURCESTATUS_RELEASED       ResourceStatus = "RELEASED"
	RESOURCESTATUS_UNCHANGED      ResourceStatus = "UNCHANGED"
	RESOURCESTATUS_TRANSFERRED    ResourceStatus = "TRANSFERRED"
	RESOURCESTATUS_UPDATED        ResourceStatus = "UPDATED"
	RESOURCESTATUS_ALT_ANCHOR_SMF ResourceStatus = "ALT_ANCHOR_SMF"
)
