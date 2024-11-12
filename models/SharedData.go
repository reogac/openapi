package models

type SharedData struct {
	SharedSmSubsData        *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty"`
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty"`
	TreatmentInstructions   map[string]string                  `json:"treatmentInstructions,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo                 `json:"sharedEcsAddrConfigInfo,omitempty"`
	SharedDataId            string                             `json:"sharedDataId"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
}
