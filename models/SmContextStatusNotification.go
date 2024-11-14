/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextStatusNotification struct {
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
}
