package models

type StatusNotification struct {
	StatusInfo          StatusInfo           `json:"statusInfo"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
}
