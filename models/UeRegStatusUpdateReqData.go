package models
type UeRegStatusUpdateReqData struct {
	 TransferStatus	string	`json:"transferStatus"`
	 ToReleaseSessionList	[]int	`json:"toReleaseSessionList,omitempty"`
	 PcfReselectedInd	*bool	`json:"pcfReselectedInd,omitempty"`
	 SmfChangeInfoList	[]SmfChangeInfo	`json:"smfChangeInfoList,omitempty"`
}
