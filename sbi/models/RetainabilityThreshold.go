/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RetainabilityThreshold struct {
	RelFlowNum   *int     `json:"relFlowNum,omitempty"`
	RelTimeUnit  TimeUnit `json:"relTimeUnit,omitempty"`
	RelFlowRatio *int     `json:"relFlowRatio,omitempty"`
}
