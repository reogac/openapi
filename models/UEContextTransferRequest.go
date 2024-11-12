package models

type UEContextTransferRequest struct {
	JsonData            *UeContextTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN1Message string                    `json:"binaryDataN1Message,omitempty"`
}
