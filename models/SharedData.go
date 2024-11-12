package models

type SharedData struct {
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty"`
	SharedSmSubsData        *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo                 `json:"sharedEcsAddrConfigInfo,omitempty"`
	SharedDataId            string                             `json:"sharedDataId"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty"`
	TreatmentInstructions   map[string]string                  `json:"treatmentInstructions,omitempty"`
}
