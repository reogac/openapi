package models

type AssignEbiError struct {
	Error          ProblemDetails  `json:"error"`
	FailureDetails AssignEbiFailed `json:"failureDetails"`
}
