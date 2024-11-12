package models

type SendMoDataRequest struct {
	JsonData     *SendMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData string             `json:"binaryMoData,omitempty"`
}
