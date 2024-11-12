package models

type N1N2MessageTransferRequest struct {
	JsonData                *N1N2MessageTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN1Message     string                      `json:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information string                      `json:"binaryDataN2Information,omitempty"`
	BinaryMtData            string                      `json:"binaryMtData,omitempty"`
}
