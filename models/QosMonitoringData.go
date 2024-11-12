package models

type QosMonitoringData struct {
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	QmId            string   `json:"qmId"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
}
