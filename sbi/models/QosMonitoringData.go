/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	QmId            string   `json:"qmId"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
}
