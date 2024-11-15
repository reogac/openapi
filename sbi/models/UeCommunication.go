/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCommunication struct {
	Confidence        *int                        `json:"confidence,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
}
