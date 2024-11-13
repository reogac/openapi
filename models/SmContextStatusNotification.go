package models

type SmContextStatusNotification struct {
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
}
