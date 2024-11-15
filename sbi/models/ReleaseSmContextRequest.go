/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleaseSmContextRequest struct {
	JsonData                  *SmContextReleaseData `json:"jsonData,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty"`
}
