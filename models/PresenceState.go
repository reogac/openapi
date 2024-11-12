package models

type PresenceState string

// Define constant values for PresenceState
const (
	PRESENCESTATE_IN_AREA     PresenceState = "IN_AREA"
	PRESENCESTATE_OUT_OF_AREA PresenceState = "OUT_OF_AREA"
	PRESENCESTATE_UNKNOWN     PresenceState = "UNKNOWN"
	PRESENCESTATE_INACTIVE    PresenceState = "INACTIVE"
)
