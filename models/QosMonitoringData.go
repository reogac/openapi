package models

type QosMonitoringData struct {
	RepFreqs        []string `json:"repFreqs"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	QmId            string   `json:"qmId"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
}
