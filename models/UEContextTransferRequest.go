package models
type UEContextTransferRequest struct {
	 JsonData	*UeContextTransferReqData	`json:"jsonData,omitempty"`
	 BinaryDataN1Message	[]byte	`json:"binaryDataN1Message,omitempty"`
}
