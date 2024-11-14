package models
type RegistrationContextContainer struct {
	 SelectedPlmnId	*PlmnId	`json:"selectedPlmnId,omitempty"`
	 LocalTimeZone	string	`json:"localTimeZone,omitempty"`
	 InitialAmfName	string	`json:"initialAmfName"`
	 UserLocation	UserLocation	`json:"userLocation"`
	 InitialAmfN2ApId	*int	`json:"initialAmfN2ApId,omitempty"`
	 AnN2IPv4Addr	string	`json:"anN2IPv4Addr,omitempty"`
	 AnN2IPv6Addr	string	`json:"anN2IPv6Addr,omitempty"`
	 AnType	AccessType	`json:"anType"`
	 AnN2ApId	int	`json:"anN2ApId"`
	 RrcEstCause	string	`json:"rrcEstCause,omitempty"`
	 IabNodeInd	*bool	`json:"iabNodeInd,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 UeContextRequest	*bool	`json:"ueContextRequest,omitempty"`
	 AllowedNssai	*AllowedNssai	`json:"allowedNssai,omitempty"`
	 ConfiguredNssai	[]ConfiguredSnssai	`json:"configuredNssai,omitempty"`
	 LteMInd	*LteMInd	`json:"lteMInd,omitempty"`
	 RanNodeId	GlobalRanNodeId	`json:"ranNodeId"`
	 RejectedNssaiInPlmn	[]Snssai	`json:"rejectedNssaiInPlmn,omitempty"`
	 RejectedNssaiInTa	[]Snssai	`json:"rejectedNssaiInTa,omitempty"`
	 CeModeBInd	*CeModeBInd	`json:"ceModeBInd,omitempty"`
	 AuthenticatedInd	*bool	`json:"authenticatedInd,omitempty"`
	 NpnAccessInfo	*NpnAccessInfo	`json:"npnAccessInfo,omitempty"`
}
