package models
type TransferMoDataRequest struct {
	 BinaryMoData	[]byte	`json:"binaryMoData,omitempty"`
	 JsonData	*TransferMoDataReqData	`json:"jsonData,omitempty"`
}
