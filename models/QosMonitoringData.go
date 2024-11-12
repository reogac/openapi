package models

type QosMonitoringData struct {
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	QmId            string   `json:"qmId"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
}
