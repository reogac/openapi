package models

type UeMobility struct {
	Ts               string                      `json:"ts,omitempty"`
	RecurringTime    *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Duration         *int                        `json:"duration,omitempty"`
	DurationVariance *float64                    `json:"durationVariance,omitempty"`
	LocInfos         []LocationInfo              `json:"locInfos,omitempty"`
}
