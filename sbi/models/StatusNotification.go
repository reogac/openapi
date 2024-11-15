/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusNotification struct {
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
}
