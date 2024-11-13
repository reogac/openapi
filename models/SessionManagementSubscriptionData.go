package models

type SessionManagementSubscriptionData struct {
	SingleNssai                     Snssai                             `json:"singleNssai"`
	DnnConfigurations               map[string]DnnConfiguration        `json:"dnnConfigurations,omitempty"`
	SharedVnGroupDataIds            map[string]string                  `json:"sharedVnGroupDataIds,omitempty"`
	SharedDnnConfigurationsId       string                             `json:"sharedDnnConfigurationsId,omitempty"`
	OdbPacketServices               OdbPacketServices                  `json:"odbPacketServices,omitempty"`
	SharedTraceDataId               string                             `json:"sharedTraceDataId,omitempty"`
	ExpectedUeBehavioursList        map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	SuggestedPacketNumDlList        map[string]SuggestedPacketNumDl    `json:"suggestedPacketNumDlList,omitempty"`
	ThreeGppChargingCharacteristics string                             `json:"3gppChargingCharacteristics,omitempty"`
	SupportedFeatures               string                             `json:"supportedFeatures,omitempty"`
	InternalGroupIds                []string                           `json:"internalGroupIds,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
}
