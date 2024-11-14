/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeRegStatusUpdateReqData struct {
	ToReleaseSessionList []int                   `json:"toReleaseSessionList,omitempty"`
	PcfReselectedInd     *bool                   `json:"pcfReselectedInd,omitempty"`
	SmfChangeInfoList    []SmfChangeInfo         `json:"smfChangeInfoList,omitempty"`
	AnalyticsNotUsedList []string                `json:"analyticsNotUsedList,omitempty"`
	ToReleaseSessionInfo []ReleaseSessionInfo    `json:"toReleaseSessionInfo,omitempty"`
	TransferStatus       UeContextTransferStatus `json:"transferStatus"`
}
