package models

type SubscriptionDataSets struct {
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
}
