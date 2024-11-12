package models

type UeContextTransferReqData struct {
	Reason            TransferReason      `json:"reason"`
	AccessType        AccessType          `json:"accessType"`
	PlmnId            *PlmnId             `json:"plmnId,omitempty"`
	RegRequest        *N1MessageContainer `json:"regRequest,omitempty"`
	SupportedFeatures string              `json:"supportedFeatures,omitempty"`
}
