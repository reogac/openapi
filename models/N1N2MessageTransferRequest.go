package models

type N1N2MessageTransferRequest struct {
	JsonData                *N1N2MessageTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN1Message     []byte                      `json:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information []byte                      `json:"binaryDataN2Information,omitempty"`
	BinaryMtData            []byte                      `json:"binaryMtData,omitempty"`
}