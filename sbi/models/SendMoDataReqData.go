/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SendMoDataReqData struct {
	UeLocation       *UserLocation     `json:"ueLocation,omitempty"`
	MoData           RefToBinaryData   `json:"moData"`
	MoExpDataCounter *MoExpDataCounter `json:"moExpDataCounter,omitempty"`
}
