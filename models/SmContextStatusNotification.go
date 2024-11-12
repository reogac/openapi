package models

type SmContextStatusNotification struct {
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
}
