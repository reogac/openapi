/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostPduSessionsErrorResponse struct {
	JsonData               *PduSessionCreateError `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoToUe []byte                 `json:"binaryDataN1SmInfoToUe,omitempty"`
}
