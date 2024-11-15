/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CancelRelocateUEContextRequest struct {
	JsonData              *UeContextCancelRelocateData `json:"jsonData,omitempty"`
	BinaryDataGtpcMessage []byte                       `json:"binaryDataGtpcMessage,omitempty"`
}
