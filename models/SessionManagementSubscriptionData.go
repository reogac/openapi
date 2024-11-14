package models
type SessionManagementSubscriptionData struct {
	 SharedTraceDataId	string	`json:"sharedTraceDataId,omitempty"`
	 SuggestedPacketNumDlList	map[string]SuggestedPacketNumDl	`json:"suggestedPacketNumDlList,omitempty"`
	 ThreeGppChargingCharacteristics	string	`json:"3gppChargingCharacteristics,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 DnnConfigurations	map[string]DnnConfiguration	`json:"dnnConfigurations,omitempty"`
	 SharedVnGroupDataIds	map[string]string	`json:"sharedVnGroupDataIds,omitempty"`
	 SharedDnnConfigurationsId	string	`json:"sharedDnnConfigurationsId,omitempty"`
	 OdbPacketServices	OdbPacketServices	`json:"odbPacketServices,omitempty"`
	 SingleNssai	Snssai	`json:"singleNssai"`
	 InternalGroupIds	[]string	`json:"internalGroupIds,omitempty"`
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 ExpectedUeBehavioursList	map[string]ExpectedUeBehaviourData	`json:"expectedUeBehavioursList,omitempty"`
}
