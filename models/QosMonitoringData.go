package models
type QosMonitoringData struct {
	 DirectNotifInd	*bool	`json:"directNotifInd,omitempty"`
	 QmId	string	`json:"qmId"`
	 ReqQosMonParams	[]string	`json:"reqQosMonParams"`
	 RepFreqs	[]string	`json:"repFreqs"`
	 RepThreshRp	*int	`json:"repThreshRp,omitempty"`
	 WaitTime	*int	`json:"waitTime,omitempty"`
	 RepThreshDl	*int	`json:"repThreshDl,omitempty"`
	 RepThreshUl	*int	`json:"repThreshUl,omitempty"`
	 RepPeriod	*int	`json:"repPeriod,omitempty"`
	 NotifyUri	string	`json:"notifyUri,omitempty"`
	 NotifyCorreId	string	`json:"notifyCorreId,omitempty"`
}
