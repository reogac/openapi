package models

type UeCommunication struct {
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
}
