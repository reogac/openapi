package models

type SubscriptionDataSets struct {
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
}
