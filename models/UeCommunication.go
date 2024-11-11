package models

type UeCommunication struct {
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
}
