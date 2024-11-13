package models

type StatusNotification struct {
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
}
