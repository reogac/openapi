/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrievedData struct {
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DlDataWaitingInd    *bool                `json:"dlDataWaitingInd,omitempty"`
	AfCoordinationInfo  *AfCoordinationInfo  `json:"afCoordinationInfo,omitempty"`
	UeEpsPdnConnection  string               `json:"ueEpsPdnConnection"`
	SmContext           *SmContext           `json:"smContext,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
}
