package models

type QosMonitoringData struct {
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
}
