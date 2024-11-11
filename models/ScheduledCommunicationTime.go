type ScheduledCommunicationTime struct {
	 TimeOfDayStart	string	`json:"timeOfDayStart,omitempty"`
	 TimeOfDayEnd	string	`json:"timeOfDayEnd,omitempty"`
	 DaysOfWeek	[]int	`json:"daysOfWeek,omitempty"`
}
