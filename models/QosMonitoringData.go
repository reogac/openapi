package models

type QosMonitoringData struct {
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
}
