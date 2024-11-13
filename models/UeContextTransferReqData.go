package models

type UeContextTransferReqData struct {
	SupportedFeatures string              `json:"supportedFeatures,omitempty"`
	Reason            TransferReason      `json:"reason"`
	AccessType        AccessType          `json:"accessType"`
	PlmnId            *PlmnId             `json:"plmnId,omitempty"`
	RegRequest        *N1MessageContainer `json:"regRequest,omitempty"`
}
