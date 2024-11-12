package models

type N1N2MessageTransferError struct {
	ErrInfo *N1N2MsgTxfrErrDetail `json:"errInfo,omitempty"`
	Error   ProblemDetails        `json:"error"`
}
