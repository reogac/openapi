package models

type AuthenticationInfo struct {
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
}
