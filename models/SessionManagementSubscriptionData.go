package models

type SessionManagementSubscriptionData struct {
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	ExpectedUeBehavioursList        map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	SupportedFeatures               string                             `json:"supportedFeatures,omitempty"`
	SingleNssai                     Snssai                             `json:"singleNssai"`
	InternalGroupIds                []string                           `json:"internalGroupIds,omitempty"`
	SharedDnnConfigurationsId       string                             `json:"sharedDnnConfigurationsId,omitempty"`
	SharedTraceDataId               string                             `json:"sharedTraceDataId,omitempty"`
	SuggestedPacketNumDlList        map[string]SuggestedPacketNumDl    `json:"suggestedPacketNumDlList,omitempty"`
	ThreeGppChargingCharacteristics string                             `json:"3gppChargingCharacteristics,omitempty"`
	DnnConfigurations               map[string]DnnConfiguration        `json:"dnnConfigurations,omitempty"`
	SharedVnGroupDataIds            map[string]string                  `json:"sharedVnGroupDataIds,omitempty"`
	OdbPacketServices               OdbPacketServices                  `json:"odbPacketServices,omitempty"`
}
