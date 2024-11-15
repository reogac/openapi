package models

import (
	"fmt"
)

func NewSimpleProblem(status int, detail string) *ProblemDetails {
	return &ProblemDetails{
		Status: &status,
		Detail: detail,
	}
}

func GetProblemDetailsCode(prob *ProblemDetails) int {
	if prob.Status != nil {
		return *prob.Status
	}
	return 500
}

func ErrorFromProblemDetails(prob *ProblemDetails) error {
	if prob.Status != nil {
		return fmt.Errorf("Code: %d, Detail: %s", *prob.Status, prob.Detail)
	}

	return fmt.Errorf("Detail: %s", prob.Detail)
}
