package models

type ProSeCapability string

// Define constant values for ProSeCapability
const (
	PROSECAPABILITY_PROSE_DD           ProSeCapability = "PROSE_DD"
	PROSECAPABILITY_PROSE_DC           ProSeCapability = "PROSE_DC"
	PROSECAPABILITY_PROSE_L2_U2N_RELAY ProSeCapability = "PROSE_L2_U2N_RELAY"
	PROSECAPABILITY_PROSE_L3_U2N_RELAY ProSeCapability = "PROSE_L3_U2N_RELAY"
	PROSECAPABILITY_PROSE_L2_REMOTE_UE ProSeCapability = "PROSE_L2_REMOTE_UE"
	PROSECAPABILITY_PROSE_L3_REMOTE_UE ProSeCapability = "PROSE_L3_REMOTE_UE"
)