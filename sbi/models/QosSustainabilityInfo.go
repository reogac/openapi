/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosSustainabilityInfo struct {
	AreaInfo      *NetworkAreaInfo        `json:"areaInfo,omitempty"`
	StartTs       string                  `json:"startTs,omitempty"`
	EndTs         string                  `json:"endTs,omitempty"`
	QosFlowRetThd *RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`
	RanUeThrouThd string                  `json:"ranUeThrouThd,omitempty"`
	Snssai        *Snssai                 `json:"snssai,omitempty"`
	Confidence    *int                    `json:"confidence,omitempty"`
}
