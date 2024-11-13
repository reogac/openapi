package models

type QosMonitoringData struct {
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	QmId            string   `json:"qmId"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
}
