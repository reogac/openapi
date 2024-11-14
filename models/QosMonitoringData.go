/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	QmId            string   `json:"qmId"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
}
