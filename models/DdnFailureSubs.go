/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DdnFailureSubs struct {
	DdnFailureSubsInd      *bool               `json:"ddnFailureSubsInd,omitempty"`
	DdnFailureSubsInfoList []DdnFailureSubInfo `json:"ddnFailureSubsInfoList,omitempty"`
}
