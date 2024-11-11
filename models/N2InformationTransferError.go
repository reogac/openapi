package models
type N2InformationTransferError struct {
	 Error	ProblemDetails	`json:"error"`
	 PwsErrorInfo	*PWSErrorData	`json:"pwsErrorInfo,omitempty"`
}
