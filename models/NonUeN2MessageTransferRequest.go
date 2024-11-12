package models

type NonUeN2MessageTransferRequest struct {
	JsonData                *N2InformationTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN2Information string                        `json:"binaryDataN2Information,omitempty"`
}
