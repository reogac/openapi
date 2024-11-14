package models
type TscaiInputContainer struct {
	 BurstArrivalTime	string	`json:"burstArrivalTime,omitempty"`
	 SurTimeInNumMsg	*int	`json:"surTimeInNumMsg,omitempty"`
	 SurTimeInTime	*int	`json:"surTimeInTime,omitempty"`
	 Periodicity	*int	`json:"periodicity,omitempty"`
}
