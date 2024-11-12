package models

type UEContextTransferRequest struct {
	BinaryDataN1Message string                    `json:"binaryDataN1Message,omitempty"`
	JsonData            *UeContextTransferReqData `json:"jsonData,omitempty"`
}
