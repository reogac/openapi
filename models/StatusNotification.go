package models

type StatusNotification struct {
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
}
