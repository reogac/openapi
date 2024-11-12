package models

type StatusNotification struct {
	StatusInfo          StatusInfo           `json:"statusInfo"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
}
