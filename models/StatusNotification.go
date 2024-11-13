package models

type StatusNotification struct {
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
}
