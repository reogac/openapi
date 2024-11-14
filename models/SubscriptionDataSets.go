package models
type SubscriptionDataSets struct {
	 SmfSelData	*SmfSelectionSubscriptionData	`json:"smfSelData,omitempty"`
	 UecAmfData	*UeContextInAmfData	`json:"uecAmfData,omitempty"`
	 UecSmfData	*UeContextInSmfData	`json:"uecSmfData,omitempty"`
	 UecSmsfData	*UeContextInSmsfData	`json:"uecSmsfData,omitempty"`
	 SmsSubsData	*SmsSubscriptionData	`json:"smsSubsData,omitempty"`
	 SmData	*SmSubsData	`json:"smData,omitempty"`
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 LcsBroadcastAssistanceTypesData	*LcsBroadcastAssistanceTypesData	`json:"lcsBroadcastAssistanceTypesData,omitempty"`
	 MbsData	*MbsSubscriptionData	`json:"mbsData,omitempty"`
	 SmsMngData	*SmsManagementSubscriptionData	`json:"smsMngData,omitempty"`
	 V2xData	*V2xSubscriptionData	`json:"v2xData,omitempty"`
	 UcData	*UcSubscriptionData	`json:"ucData,omitempty"`
	 AmData	*AccessAndMobilitySubscriptionData	`json:"amData,omitempty"`
	 LcsPrivacyData	*LcsPrivacyData	`json:"lcsPrivacyData,omitempty"`
	 LcsMoData	*LcsMoData	`json:"lcsMoData,omitempty"`
	 ProseData	*ProseSubscriptionData	`json:"proseData,omitempty"`
}
