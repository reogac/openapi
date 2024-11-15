/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateSmContextRequest struct {
	BinaryDataN1SmMessage         []byte               `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation     []byte               `json:"binaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationExt1 []byte               `json:"binaryDataN2SmInformationExt1,omitempty"`
	JsonData                      *SmContextUpdateData `json:"jsonData,omitempty"`
}
