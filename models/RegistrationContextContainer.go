package models
type RegistrationContextContainer struct {
	 LocalTimeZone	string	`json:"localTimeZone,omitempty"`
	 UserLocation	UserLocation	`json:"userLocation"`
	 AnN2IPv4Addr	string	`json:"anN2IPv4Addr,omitempty"`
	 AllowedNssai	*AllowedNssai	`json:"allowedNssai,omitempty"`
	 RejectedNssaiInTa	[]Snssai	`json:"rejectedNssaiInTa,omitempty"`
	 SelectedPlmnId	*PlmnId	`json:"selectedPlmnId,omitempty"`
	 NpnAccessInfo	*NpnAccessInfo	`json:"npnAccessInfo,omitempty"`
	 InitialAmfName	string	`json:"initialAmfName"`
	 RrcEstCause	string	`json:"rrcEstCause,omitempty"`
	 UeContextRequest	*bool	`json:"ueContextRequest,omitempty"`
	 RejectedNssaiInPlmn	[]Snssai	`json:"rejectedNssaiInPlmn,omitempty"`
	 IabNodeInd	*bool	`json:"iabNodeInd,omitempty"`
	 LteMInd	*LteMInd	`json:"lteMInd,omitempty"`
	 AnN2ApId	int	`json:"anN2ApId"`
	 RanNodeId	GlobalRanNodeId	`json:"ranNodeId"`
	 ConfiguredNssai	[]ConfiguredSnssai	`json:"configuredNssai,omitempty"`
	 AuthenticatedInd	*bool	`json:"authenticatedInd,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 AnType	AccessType	`json:"anType"`
	 InitialAmfN2ApId	*int	`json:"initialAmfN2ApId,omitempty"`
	 AnN2IPv6Addr	string	`json:"anN2IPv6Addr,omitempty"`
	 CeModeBInd	*CeModeBInd	`json:"ceModeBInd,omitempty"`
}
