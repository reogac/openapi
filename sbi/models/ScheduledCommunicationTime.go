/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ScheduledCommunicationTime struct {
	TimeOfDayEnd   string `json:"timeOfDayEnd,omitempty"`
	DaysOfWeek     []int  `json:"daysOfWeek,omitempty"`
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`
}
