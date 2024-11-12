package models

type TscaiInputContainer struct {
	Periodicity      *int   `json:"periodicity,omitempty"`
	BurstArrivalTime string `json:"burstArrivalTime,omitempty"`
	SurTimeInNumMsg  *int   `json:"surTimeInNumMsg,omitempty"`
	SurTimeInTime    *int   `json:"surTimeInTime,omitempty"`
}
