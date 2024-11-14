package models
type DnnConfiguration struct {
	 OnboardingInd	*bool	`json:"onboardingInd,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 StaticIpAddress	[]IpAddress	`json:"staticIpAddress,omitempty"`
	 NiddInfo	*NiddInformation	`json:"niddInfo,omitempty"`
	 SecondaryAuth	*bool	`json:"secondaryAuth,omitempty"`
	 AdditionalSharedEcsAddrConfigInfoIds	[]string	`json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	 PduSessionContinuityInd	PduSessionContinuityInd	`json:"pduSessionContinuityInd,omitempty"`
	 Ipv6FrameRouteList	[]FrameRouteInfo	`json:"ipv6FrameRouteList,omitempty"`
	 DnAaaFqdn	string	`json:"dnAaaFqdn,omitempty"`
	 Ipv6Index	*IpIndex	`json:"ipv6Index,omitempty"`
	 RedundantSessionAllowed	*bool	`json:"redundantSessionAllowed,omitempty"`
	 AcsInfo	*AcsInfo	`json:"acsInfo,omitempty"`
	 Ipv4FrameRouteList	[]FrameRouteInfo	`json:"ipv4FrameRouteList,omitempty"`
	 Ipv4Index	*IpIndex	`json:"ipv4Index,omitempty"`
	 PduSessionTypes	PduSessionTypes	`json:"pduSessionTypes"`
	 DnAaaAddress	*IpAddress	`json:"dnAaaAddress,omitempty"`
	 AdditionalDnAaaAddresses	[]IpAddress	`json:"additionalDnAaaAddresses,omitempty"`
	 DnAaaIpAddressAllocation	*bool	`json:"dnAaaIpAddressAllocation,omitempty"`
	 AerialUeInd	AerialUeIndication	`json:"aerialUeInd,omitempty"`
	 SscModes	SscModes	`json:"sscModes"`
	 FiveGQosProfile	*SubscribedDefaultQos	`json:"5gQosProfile,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 NiddNefId	string	`json:"niddNefId,omitempty"`
	 ThreeGppChargingCharacteristics	string	`json:"3gppChargingCharacteristics,omitempty"`
	 AtsssAllowed	*bool	`json:"atsssAllowed,omitempty"`
	 SharedEcsAddrConfigInfo	string	`json:"sharedEcsAddrConfigInfo,omitempty"`
	 EasDiscoveryAuthorized	*bool	`json:"easDiscoveryAuthorized,omitempty"`
	 IwkEpsInd	*bool	`json:"iwkEpsInd,omitempty"`
	 IptvAccCtrlInfo	string	`json:"iptvAccCtrlInfo,omitempty"`
	 EcsAddrConfigInfo	*EcsAddrConfigInfo	`json:"ecsAddrConfigInfo,omitempty"`
	 SubscribedMaxIpv6PrefixSize	*int	`json:"subscribedMaxIpv6PrefixSize,omitempty"`
	 UavSecondaryAuth	*bool	`json:"uavSecondaryAuth,omitempty"`
	 AdditionalEcsAddrConfigInfos	[]EcsAddrConfigInfo	`json:"additionalEcsAddrConfigInfos,omitempty"`
}
