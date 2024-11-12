package models

type TransferMoDataRequest struct {
	BinaryMoData string                 `json:"binaryMoData,omitempty"`
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
}
