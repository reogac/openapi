package models

type AuthenticationInfo struct {
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
}
