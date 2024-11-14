package models
type UeCommunication struct {
	 CommDur	*int	`json:"commDur,omitempty"`
	 Ts	string	`json:"ts,omitempty"`
	 AnaOfAppList	*AppListForUeComm	`json:"anaOfAppList,omitempty"`
	 TsVariance	*float64	`json:"tsVariance,omitempty"`
	 RecurringTime	*ScheduledCommunicationTime	`json:"recurringTime,omitempty"`
	 TrafChar	*TrafficCharacterization	`json:"trafChar,omitempty"`
	 Ratio	*int	`json:"ratio,omitempty"`
	 PerioCommInd	*bool	`json:"perioCommInd,omitempty"`
	 CommDurVariance	*float64	`json:"commDurVariance,omitempty"`
	 PerioTime	*int	`json:"perioTime,omitempty"`
	 PerioTimeVariance	*float64	`json:"perioTimeVariance,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 SessInactTimer	*SessInactTimerForUeComm	`json:"sessInactTimer,omitempty"`
}
