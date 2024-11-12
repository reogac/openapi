package models

type UeCommunication struct {
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
}
