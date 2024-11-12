package models

type FailureEventInfo struct {
	FailureCode NwdafFailureCode `json:"failureCode"`
	Event       NwdafEvent       `json:"event"`
}
