/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SharedData struct {
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty"`
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo                 `json:"sharedEcsAddrConfigInfo,omitempty"`
	SharedDataId            string                             `json:"sharedDataId"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty"`
	TreatmentInstructions   map[string]string                  `json:"treatmentInstructions,omitempty"`
	SharedSmSubsData        *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
}
