/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
}
