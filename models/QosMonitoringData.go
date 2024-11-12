package models

type QosMonitoringData struct {
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
}
