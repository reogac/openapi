package models
type ScheduledCommunicationTime struct {
	 DaysOfWeek	[]int	`json:"daysOfWeek,omitempty"`
	 TimeOfDayStart	string	`json:"timeOfDayStart,omitempty"`
	 TimeOfDayEnd	string	`json:"timeOfDayEnd,omitempty"`
}
