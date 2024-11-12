package models

type SessionManagementSubscriptionData struct {
	DnnConfigurations               map[string]DnnConfiguration        `json:"dnnConfigurations,omitempty"`
	SharedTraceDataId               string                             `json:"sharedTraceDataId,omitempty"`
	ExpectedUeBehavioursList        map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	SupportedFeatures               string                             `json:"supportedFeatures,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	SuggestedPacketNumDlList        map[string]SuggestedPacketNumDl    `json:"suggestedPacketNumDlList,omitempty"`
	ThreeGppChargingCharacteristics string                             `json:"3gppChargingCharacteristics,omitempty"`
	SingleNssai                     Snssai                             `json:"singleNssai"`
	InternalGroupIds                []string                           `json:"internalGroupIds,omitempty"`
	SharedVnGroupDataIds            map[string]string                  `json:"sharedVnGroupDataIds,omitempty"`
	SharedDnnConfigurationsId       string                             `json:"sharedDnnConfigurationsId,omitempty"`
	OdbPacketServices               OdbPacketServices                  `json:"odbPacketServices,omitempty"`
}
