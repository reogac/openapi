package models

type AuthenticationInfo struct {
	N5gcInd               *bool                  `json:"n5gcInd,omitempty"`
	OnboardingInd         *bool                  `json:"onboardingInd,omitempty"`
	SupiOrSuci            string                 `json:"supiOrSuci"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	Pei                   string                 `json:"pei,omitempty"`
	TraceData             *TraceData             `json:"traceData,omitempty"`
	RoutingIndicator      string                 `json:"routingIndicator,omitempty"`
	DisasterRoamingInd    *bool                  `json:"disasterRoamingInd,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	UdmGroupId            string                 `json:"udmGroupId,omitempty"`
	CellCagInfo           []string               `json:"cellCagInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	NswoInd               *bool                  `json:"nswoInd,omitempty"`
}
