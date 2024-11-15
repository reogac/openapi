/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TransferMoDataRequest struct {
	BinaryMoData []byte                 `json:"binaryMoData,omitempty"`
	JsonData     *TransferMoDataReqData `json:"jsonData,omitempty"`
}
