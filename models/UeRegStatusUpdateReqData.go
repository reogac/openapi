package models
type UeRegStatusUpdateReqData struct {
	 ToReleaseSessionList	[]int	`json:"toReleaseSessionList,omitempty"`
	 PcfReselectedInd	*bool	`json:"pcfReselectedInd,omitempty"`
	 SmfChangeInfoList	[]SmfChangeInfo	`json:"smfChangeInfoList,omitempty"`
	 AnalyticsNotUsedList	[]string	`json:"analyticsNotUsedList,omitempty"`
	 ToReleaseSessionInfo	[]ReleaseSessionInfo	`json:"toReleaseSessionInfo,omitempty"`
	 TransferStatus	UeContextTransferStatus	`json:"transferStatus"`
}
