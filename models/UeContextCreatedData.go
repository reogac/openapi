package models
type UeContextCreatedData struct {
	 TargetToSourceData	N2InfoContent	`json:"targetToSourceData"`
	 PduSessionList	[]N2SmInformation	`json:"pduSessionList"`
	 FailedSessionList	[]N2SmInformation	`json:"failedSessionList,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 PcfReselectedInd	*bool	`json:"pcfReselectedInd,omitempty"`
	 AnalyticsNotUsedList	[]string	`json:"analyticsNotUsedList,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
}
