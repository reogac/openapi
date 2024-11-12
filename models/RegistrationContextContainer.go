package models
type RegistrationContextContainer struct {
	 IabNodeInd	*bool	`json:"iabNodeInd,omitempty"`
	 LteMInd	*LteMInd	`json:"lteMInd,omitempty"`
	 LocalTimeZone	string	`json:"localTimeZone,omitempty"`
	 AnN2ApId	int	`json:"anN2ApId"`
	 AnN2IPv4Addr	string	`json:"anN2IPv4Addr,omitempty"`
	 RejectedNssaiInTa	[]Snssai	`json:"rejectedNssaiInTa,omitempty"`
	 AnType	AccessType	`json:"anType"`
	 InitialAmfName	string	`json:"initialAmfName"`
	 NpnAccessInfo	*NpnAccessInfo	`json:"npnAccessInfo,omitempty"`
	 AllowedNssai	*AllowedNssai	`json:"allowedNssai,omitempty"`
	 SelectedPlmnId	*PlmnId	`json:"selectedPlmnId,omitempty"`
	 CeModeBInd	*CeModeBInd	`json:"ceModeBInd,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 RanNodeId	GlobalRanNodeId	`json:"ranNodeId"`
	 RrcEstCause	string	`json:"rrcEstCause,omitempty"`
	 AnN2IPv6Addr	string	`json:"anN2IPv6Addr,omitempty"`
	 RejectedNssaiInPlmn	[]Snssai	`json:"rejectedNssaiInPlmn,omitempty"`
	 AuthenticatedInd	*bool	`json:"authenticatedInd,omitempty"`
	 UserLocation	UserLocation	`json:"userLocation"`
	 UeContextRequest	*bool	`json:"ueContextRequest,omitempty"`
	 InitialAmfN2ApId	*int	`json:"initialAmfN2ApId,omitempty"`
	 ConfiguredNssai	[]ConfiguredSnssai	`json:"configuredNssai,omitempty"`
}
