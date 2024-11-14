/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMoDataRequest struct {
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData []byte                 `json:"binaryMoData,omitempty"`
}
