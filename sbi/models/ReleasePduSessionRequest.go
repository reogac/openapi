/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleasePduSessionRequest struct {
	BinaryDataN4InformationExt1 []byte       `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 []byte       `json:"binaryDataN4InformationExt2,omitempty"`
	JsonData                    *ReleaseData `json:"jsonData,omitempty"`
	BinaryDataN4Information     []byte       `json:"binaryDataN4Information,omitempty"`
}
