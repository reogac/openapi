package models

type SmContextStatusNotification struct {
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
}
