package models

type UeCommunication struct {
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
}
