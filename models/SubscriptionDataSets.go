package models

type SubscriptionDataSets struct {
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
}
