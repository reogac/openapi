package models

type UeCommunication struct {
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
}
