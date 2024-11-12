package models

type SharedData struct {
	SharedSmSubsData        *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty"`
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty"`
	SharedDataId            string                             `json:"sharedDataId"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty"`
	TreatmentInstructions   map[string]string                  `json:"treatmentInstructions,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo                 `json:"sharedEcsAddrConfigInfo,omitempty"`
}
