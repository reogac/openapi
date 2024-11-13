package models

type UeRegStatusUpdateReqData struct {
	SmfChangeInfoList    []SmfChangeInfo         `json:"smfChangeInfoList,omitempty"`
	AnalyticsNotUsedList []string                `json:"analyticsNotUsedList,omitempty"`
	ToReleaseSessionInfo []ReleaseSessionInfo    `json:"toReleaseSessionInfo,omitempty"`
	TransferStatus       UeContextTransferStatus `json:"transferStatus"`
	ToReleaseSessionList []int                   `json:"toReleaseSessionList,omitempty"`
	PcfReselectedInd     *bool                   `json:"pcfReselectedInd,omitempty"`
}
