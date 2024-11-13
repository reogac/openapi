package models

type SmContextStatusNotification struct {
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
}
