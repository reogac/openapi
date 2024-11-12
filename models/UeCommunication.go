package models

type UeCommunication struct {
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
}
