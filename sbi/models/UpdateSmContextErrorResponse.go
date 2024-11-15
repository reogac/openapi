/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateSmContextErrorResponse struct {
	JsonData                  *SmContextUpdateError `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage     []byte                `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
}
