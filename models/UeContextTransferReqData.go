package models

type UeContextTransferReqData struct {
	Reason            string              `json:"reason"`
	AccessType        string              `json:"accessType"`
	PlmnId            *PlmnId             `json:"plmnId,omitempty"`
	RegRequest        *N1MessageContainer `json:"regRequest,omitempty"`
	SupportedFeatures string              `json:"supportedFeatures,omitempty"`
}
