package models
type SharedData struct {
	 SharedAmData	*AccessAndMobilitySubscriptionData	`json:"sharedAmData,omitempty"`
	 SharedSmsMngSubsData	*SmsManagementSubscriptionData	`json:"sharedSmsMngSubsData,omitempty"`
	 SharedSnssaiInfos	map[string]SnssaiInfo	`json:"sharedSnssaiInfos,omitempty"`
	 SharedVnGroupDatas	map[string]VnGroupData	`json:"sharedVnGroupDatas,omitempty"`
	 SharedSmSubsData	*SessionManagementSubscriptionData	`json:"sharedSmSubsData,omitempty"`
	 SharedDataId	string	`json:"sharedDataId"`
	 SharedSmsSubsData	*SmsSubscriptionData	`json:"sharedSmsSubsData,omitempty"`
	 SharedDnnConfigurations	map[string]DnnConfiguration	`json:"sharedDnnConfigurations,omitempty"`
	 SharedTraceData	*TraceData	`json:"sharedTraceData,omitempty"`
	 TreatmentInstructions	map[string]string	`json:"treatmentInstructions,omitempty"`
	 SharedEcsAddrConfigInfo	*EcsAddrConfigInfo	`json:"sharedEcsAddrConfigInfo,omitempty"`
}
