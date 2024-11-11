package models
type SmContextRetrievedData struct {
	 UeEpsPdnConnection	string	`json:"ueEpsPdnConnection"`
	 SmContext	*SmContext	`json:"smContext,omitempty"`
	 SmallDataRateStatus	*SmallDataRateStatus	`json:"smallDataRateStatus,omitempty"`
	 ApnRateStatus	*ApnRateStatus	`json:"apnRateStatus,omitempty"`
	 DlDataWaitingInd	*bool	`json:"dlDataWaitingInd,omitempty"`
}
