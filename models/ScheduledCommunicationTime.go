/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ScheduledCommunicationTime struct {
	TimeOfDayEnd   string `json:"timeOfDayEnd,omitempty"`
	DaysOfWeek     []int  `json:"daysOfWeek,omitempty"`
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`
}
