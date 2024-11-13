package models

type SendMoDataRequest struct {
	JsonData     *SendMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData []byte             `json:"binaryMoData,omitempty"`
}
