/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSets struct {
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
}
