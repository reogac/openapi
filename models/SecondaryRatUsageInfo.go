/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecondaryRatUsageInfo struct {
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
	SecondaryRatType    RatType              `json:"secondaryRatType"`
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
}
