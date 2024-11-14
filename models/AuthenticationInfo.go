package models
type AuthenticationInfo struct {
	 TraceData	*TraceData	`json:"traceData,omitempty"`
	 UdmGroupId	string	`json:"udmGroupId,omitempty"`
	 NswoInd	*bool	`json:"nswoInd,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 OnboardingInd	*bool	`json:"onboardingInd,omitempty"`
	 ServingNetworkName	string	`json:"servingNetworkName"`
	 ResynchronizationInfo	*ResynchronizationInfo	`json:"resynchronizationInfo,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 N5gcInd	*bool	`json:"n5gcInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SupiOrSuci	string	`json:"supiOrSuci"`
	 RoutingIndicator	string	`json:"routingIndicator,omitempty"`
	 CellCagInfo	[]string	`json:"cellCagInfo,omitempty"`
}
