/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEContextTransferResponse struct {
	BinaryDataN2Information     []byte                    `json:"binaryDataN2Information,omitempty"`
	BinaryDataN2InformationExt1 []byte                    `json:"binaryDataN2InformationExt1,omitempty"`
	BinaryDataN2InformationExt2 []byte                    `json:"binaryDataN2InformationExt2,omitempty"`
	JsonData                    *UeContextTransferRspData `json:"jsonData,omitempty"`
}
