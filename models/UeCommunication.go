package models
type UeCommunication struct {
	 CommDurVariance	*float64	`json:"commDurVariance,omitempty"`
	 TsVariance	*float64	`json:"tsVariance,omitempty"`
	 TrafChar	*TrafficCharacterization	`json:"trafChar,omitempty"`
	 Ratio	*int	`json:"ratio,omitempty"`
	 PerioCommInd	*bool	`json:"perioCommInd,omitempty"`
	 AnaOfAppList	*AppListForUeComm	`json:"anaOfAppList,omitempty"`
	 CommDur	*int	`json:"commDur,omitempty"`
	 PerioTime	*int	`json:"perioTime,omitempty"`
	 PerioTimeVariance	*float64	`json:"perioTimeVariance,omitempty"`
	 Ts	string	`json:"ts,omitempty"`
	 RecurringTime	*ScheduledCommunicationTime	`json:"recurringTime,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 SessInactTimer	*SessInactTimerForUeComm	`json:"sessInactTimer,omitempty"`
}
