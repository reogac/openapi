package models

type ScheduledCommunicationTime struct {
	TimeOfDayEnd   string `json:"timeOfDayEnd,omitempty"`
	DaysOfWeek     []int  `json:"daysOfWeek,omitempty"`
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`
}
