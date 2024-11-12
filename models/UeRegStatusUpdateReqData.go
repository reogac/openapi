package models

type UeRegStatusUpdateReqData struct {
	TransferStatus       UeContextTransferStatus `json:"transferStatus"`
	ToReleaseSessionList []int                   `json:"toReleaseSessionList,omitempty"`
	PcfReselectedInd     *bool                   `json:"pcfReselectedInd,omitempty"`
	SmfChangeInfoList    []SmfChangeInfo         `json:"smfChangeInfoList,omitempty"`
	AnalyticsNotUsedList []string                `json:"analyticsNotUsedList,omitempty"`
	ToReleaseSessionInfo []ReleaseSessionInfo    `json:"toReleaseSessionInfo,omitempty"`
}
