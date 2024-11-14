package models
type QosMonitoringData struct {
	 RepThreshRp	*int	`json:"repThreshRp,omitempty"`
	 WaitTime	*int	`json:"waitTime,omitempty"`
	 NotifyUri	string	`json:"notifyUri,omitempty"`
	 NotifyCorreId	string	`json:"notifyCorreId,omitempty"`
	 QmId	string	`json:"qmId"`
	 ReqQosMonParams	[]string	`json:"reqQosMonParams"`
	 RepThreshUl	*int	`json:"repThreshUl,omitempty"`
	 DirectNotifInd	*bool	`json:"directNotifInd,omitempty"`
	 RepFreqs	[]string	`json:"repFreqs"`
	 RepThreshDl	*int	`json:"repThreshDl,omitempty"`
	 RepPeriod	*int	`json:"repPeriod,omitempty"`
}
