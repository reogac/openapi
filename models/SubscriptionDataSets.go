package models

type SubscriptionDataSets struct {
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
}
