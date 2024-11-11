package models

type UeContextTransferReqData struct {
	AccessType        string              `json:"accessType"`
	PlmnId            *PlmnId             `json:"plmnId,omitempty"`
	RegRequest        *N1MessageContainer `json:"regRequest,omitempty"`
	SupportedFeatures string              `json:"supportedFeatures,omitempty"`
	Reason            string              `json:"reason"`
}
