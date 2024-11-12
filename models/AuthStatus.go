package models

type AuthStatus string

// Define constant values for AuthStatus
const (
	AUTHSTATUS_EAP_SUCCESS AuthStatus = "EAP_SUCCESS"
	AUTHSTATUS_EAP_FAILURE AuthStatus = "EAP_FAILURE"
	AUTHSTATUS_PENDING     AuthStatus = "PENDING"
)
