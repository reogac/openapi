package models

type PduSessionRelCause string

// Define constant values for PduSessionRelCause
const (
	PDUSESSIONRELCAUSE_PS_TO_CS_HO PduSessionRelCause = "PS_TO_CS_HO"
	PDUSESSIONRELCAUSE_RULE_ERROR  PduSessionRelCause = "RULE_ERROR"
)
