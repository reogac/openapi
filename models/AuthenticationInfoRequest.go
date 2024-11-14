package models
type AuthenticationInfoRequest struct {
	 N5gcInd	*bool	`json:"n5gcInd,omitempty"`
	 NswoInd	*bool	`json:"nswoInd,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 ServingNetworkName	string	`json:"servingNetworkName"`
	 ResynchronizationInfo	*ResynchronizationInfo	`json:"resynchronizationInfo,omitempty"`
	 AusfInstanceId	string	`json:"ausfInstanceId"`
	 CellCagInfo	[]string	`json:"cellCagInfo,omitempty"`
}
