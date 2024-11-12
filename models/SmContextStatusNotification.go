package models

type SmContextStatusNotification struct {
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
}
