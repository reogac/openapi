package models
type TscaiInputContainer struct {
	 SurTimeInTime	*int	`json:"surTimeInTime,omitempty"`
	 Periodicity	*int	`json:"periodicity,omitempty"`
	 BurstArrivalTime	string	`json:"burstArrivalTime,omitempty"`
	 SurTimeInNumMsg	*int	`json:"surTimeInNumMsg,omitempty"`
}
