package models

type UeMobility struct {
	LocInfos         []LocationInfo              `json:"locInfos,omitempty"`
	Ts               string                      `json:"ts,omitempty"`
	RecurringTime    *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Duration         *int                        `json:"duration,omitempty"`
	DurationVariance *float64                    `json:"durationVariance,omitempty"`
}
