package models

type SmContextStatusNotification struct {
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
}
