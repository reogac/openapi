package models

type UeCommunication struct {
	Ratio             *int                        `json:"ratio,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
}
