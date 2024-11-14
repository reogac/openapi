package models
type AssignEbiError struct {
	 FailureDetails	AssignEbiFailed	`json:"failureDetails"`
	 Error	ProblemDetails	`json:"error"`
}
