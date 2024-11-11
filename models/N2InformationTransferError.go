package models

type N2InformationTransferError struct {
	PwsErrorInfo *PWSErrorData  `json:"pwsErrorInfo,omitempty"`
	Error        ProblemDetails `json:"error"`
}
