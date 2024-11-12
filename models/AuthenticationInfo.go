package models

type AuthenticationInfo struct {
	SupiOrSuci            string                 `json:"supiOrSuci"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
}
