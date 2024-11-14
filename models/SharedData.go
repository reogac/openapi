package models
type SharedData struct {
	 SharedTraceData	*TraceData	`json:"sharedTraceData,omitempty"`
	 SharedSnssaiInfos	map[string]SnssaiInfo	`json:"sharedSnssaiInfos,omitempty"`
	 SharedVnGroupDatas	map[string]VnGroupData	`json:"sharedVnGroupDatas,omitempty"`
	 TreatmentInstructions	map[string]string	`json:"treatmentInstructions,omitempty"`
	 SharedSmSubsData	*SessionManagementSubscriptionData	`json:"sharedSmSubsData,omitempty"`
	 SharedEcsAddrConfigInfo	*EcsAddrConfigInfo	`json:"sharedEcsAddrConfigInfo,omitempty"`
	 SharedAmData	*AccessAndMobilitySubscriptionData	`json:"sharedAmData,omitempty"`
	 SharedDnnConfigurations	map[string]DnnConfiguration	`json:"sharedDnnConfigurations,omitempty"`
	 SharedSmsMngSubsData	*SmsManagementSubscriptionData	`json:"sharedSmsMngSubsData,omitempty"`
	 SharedDataId	string	`json:"sharedDataId"`
	 SharedSmsSubsData	*SmsSubscriptionData	`json:"sharedSmsSubsData,omitempty"`
}
