/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeMobility struct {
	RecurringTime    *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Duration         *int                        `json:"duration,omitempty"`
	DurationVariance *float64                    `json:"durationVariance,omitempty"`
	LocInfos         []LocationInfo              `json:"locInfos,omitempty"`
	Ts               string                      `json:"ts,omitempty"`
}