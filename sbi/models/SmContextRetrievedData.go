/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrievedData struct {
	AfCoordinationInfo  *AfCoordinationInfo  `json:"afCoordinationInfo,omitempty"`
	UeEpsPdnConnection  string               `json:"ueEpsPdnConnection"`
	SmContext           *SmContext           `json:"smContext,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DlDataWaitingInd    *bool                `json:"dlDataWaitingInd,omitempty"`
}
