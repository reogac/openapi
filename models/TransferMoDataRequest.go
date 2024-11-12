package models

type TransferMoDataRequest struct {
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData string                 `json:"binaryMoData,omitempty"`
}
