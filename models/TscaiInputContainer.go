package models

type TscaiInputContainer struct {
	SurTimeInNumMsg  *int   `json:"surTimeInNumMsg,omitempty"`
	SurTimeInTime    *int   `json:"surTimeInTime,omitempty"`
	Periodicity      *int   `json:"periodicity,omitempty"`
	BurstArrivalTime string `json:"burstArrivalTime,omitempty"`
}
