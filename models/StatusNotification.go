package models

type StatusNotification struct {
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
}
