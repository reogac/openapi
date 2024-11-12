package models
type AmfEventMode struct {
	 Expiry	string	`json:"expiry,omitempty"`
	 RepPeriod	*int	`json:"repPeriod,omitempty"`
	 SampRatio	*int	`json:"sampRatio,omitempty"`
	 PartitioningCriteria	[]string	`json:"partitioningCriteria,omitempty"`
	 NotifFlag	NotificationFlag	`json:"notifFlag,omitempty"`
	 Trigger	AmfEventTrigger	`json:"trigger"`
	 MaxReports	*int	`json:"maxReports,omitempty"`
}
