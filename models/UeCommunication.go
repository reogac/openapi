package models

type UeCommunication struct {
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
}
