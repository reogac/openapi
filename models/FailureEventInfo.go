package models

type FailureEventInfo struct {
	Event       string `json:"event"`
	FailureCode string `json:"failureCode"`
}
