package models
type SubscriptionDataSets struct {
	 UecSmsfData	*UeContextInSmsfData	`json:"uecSmsfData,omitempty"`
	 SmsMngData	*SmsManagementSubscriptionData	`json:"smsMngData,omitempty"`
	 MbsData	*MbsSubscriptionData	`json:"mbsData,omitempty"`
	 SmfSelData	*SmfSelectionSubscriptionData	`json:"smfSelData,omitempty"`
	 SmsSubsData	*SmsSubscriptionData	`json:"smsSubsData,omitempty"`
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 ProseData	*ProseSubscriptionData	`json:"proseData,omitempty"`
	 AmData	*AccessAndMobilitySubscriptionData	`json:"amData,omitempty"`
	 SmData	*SmSubsData	`json:"smData,omitempty"`
	 V2xData	*V2xSubscriptionData	`json:"v2xData,omitempty"`
	 LcsBroadcastAssistanceTypesData	*LcsBroadcastAssistanceTypesData	`json:"lcsBroadcastAssistanceTypesData,omitempty"`
	 UcData	*UcSubscriptionData	`json:"ucData,omitempty"`
	 UecAmfData	*UeContextInAmfData	`json:"uecAmfData,omitempty"`
	 UecSmfData	*UeContextInSmfData	`json:"uecSmfData,omitempty"`
	 LcsPrivacyData	*LcsPrivacyData	`json:"lcsPrivacyData,omitempty"`
	 LcsMoData	*LcsMoData	`json:"lcsMoData,omitempty"`
}
