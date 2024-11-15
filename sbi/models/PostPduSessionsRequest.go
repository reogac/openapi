/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PostPduSessionsRequest struct {
	BinaryDataN1SmInfoFromUe  []byte                `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo []byte                `json:"binaryDataUnknownN1SmInfo,omitempty"`
	JsonData                  *PduSessionCreateData `json:"jsonData,omitempty"`
}
