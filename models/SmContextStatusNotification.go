package models

type SmContextStatusNotification struct {
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
}
