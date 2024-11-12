package models

type AuthenticationInfo struct {
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	Pei                   string                 `json:"pei,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
}
