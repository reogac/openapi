package models

type CreditManagementStatus string

// Define constant values for CreditManagementStatus
const (
	CREDITMANAGEMENTSTATUS_END_USER_SER_DENIED CreditManagementStatus = "END_USER_SER_DENIED"
	CREDITMANAGEMENTSTATUS_CREDIT_CTRL_NOT_APP CreditManagementStatus = "CREDIT_CTRL_NOT_APP"
	CREDITMANAGEMENTSTATUS_AUTH_REJECTED       CreditManagementStatus = "AUTH_REJECTED"
	CREDITMANAGEMENTSTATUS_USER_UNKNOWN        CreditManagementStatus = "USER_UNKNOWN"
	CREDITMANAGEMENTSTATUS_RATING_FAILED       CreditManagementStatus = "RATING_FAILED"
)