package models

type FailureEventInfo struct {
	Event       NwdafEvent       `json:"event"`
	FailureCode NwdafFailureCode `json:"failureCode"`
}
