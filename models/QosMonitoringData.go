/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	QmId            string   `json:"qmId"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
}
