package models

type ChangeType string

// Define constant values for ChangeType
const (
	CHANGETYPE_ADD     ChangeType = "ADD"
	CHANGETYPE_MOVE    ChangeType = "MOVE"
	CHANGETYPE_REMOVE  ChangeType = "REMOVE"
	CHANGETYPE_REPLACE ChangeType = "REPLACE"
)
