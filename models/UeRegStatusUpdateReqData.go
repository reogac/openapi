package models

type UeRegStatusUpdateReqData struct {
	AnalyticsNotUsedList []string             `json:"analyticsNotUsedList,omitempty"`
	ToReleaseSessionInfo []ReleaseSessionInfo `json:"toReleaseSessionInfo,omitempty"`
	TransferStatus       string               `json:"transferStatus"`
	ToReleaseSessionList []int                `json:"toReleaseSessionList,omitempty"`
	PcfReselectedInd     *bool                `json:"pcfReselectedInd,omitempty"`
	SmfChangeInfoList    []SmfChangeInfo      `json:"smfChangeInfoList,omitempty"`
}