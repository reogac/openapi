package models
type UEContextTransferRequest struct {
	 BinaryDataN1Message	[]byte	`json:"binaryDataN1Message,omitempty"`
	 JsonData	*UeContextTransferReqData	`json:"jsonData,omitempty"`
}
