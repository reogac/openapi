package models
type DnnConfiguration struct {
	 IwkEpsInd	*bool	`json:"iwkEpsInd,omitempty"`
	 UavSecondaryAuth	*bool	`json:"uavSecondaryAuth,omitempty"`
	 SscModes	SscModes	`json:"sscModes"`
	 StaticIpAddress	[]IpAddress	`json:"staticIpAddress,omitempty"`
	 Ipv4FrameRouteList	[]FrameRouteInfo	`json:"ipv4FrameRouteList,omitempty"`
	 DnAaaFqdn	string	`json:"dnAaaFqdn,omitempty"`
	 AdditionalSharedEcsAddrConfigInfoIds	[]string	`json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 NiddNefId	string	`json:"niddNefId,omitempty"`
	 Ipv6FrameRouteList	[]FrameRouteInfo	`json:"ipv6FrameRouteList,omitempty"`
	 SubscribedMaxIpv6PrefixSize	*int	`json:"subscribedMaxIpv6PrefixSize,omitempty"`
	 AdditionalEcsAddrConfigInfos	[]EcsAddrConfigInfo	`json:"additionalEcsAddrConfigInfos,omitempty"`
	 SharedEcsAddrConfigInfo	string	`json:"sharedEcsAddrConfigInfo,omitempty"`
	 OnboardingInd	*bool	`json:"onboardingInd,omitempty"`
	 NiddInfo	*NiddInformation	`json:"niddInfo,omitempty"`
	 RedundantSessionAllowed	*bool	`json:"redundantSessionAllowed,omitempty"`
	 DnAaaAddress	*IpAddress	`json:"dnAaaAddress,omitempty"`
	 Ipv6Index	*IpIndex	`json:"ipv6Index,omitempty"`
	 EcsAddrConfigInfo	*EcsAddrConfigInfo	`json:"ecsAddrConfigInfo,omitempty"`
	 FiveGQosProfile	*SubscribedDefaultQos	`json:"5gQosProfile,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 AtsssAllowed	*bool	`json:"atsssAllowed,omitempty"`
	 IptvAccCtrlInfo	string	`json:"iptvAccCtrlInfo,omitempty"`
	 PduSessionTypes	PduSessionTypes	`json:"pduSessionTypes"`
	 Ipv4Index	*IpIndex	`json:"ipv4Index,omitempty"`
	 AerialUeInd	AerialUeIndication	`json:"aerialUeInd,omitempty"`
	 EasDiscoveryAuthorized	*bool	`json:"easDiscoveryAuthorized,omitempty"`
	 AdditionalDnAaaAddresses	[]IpAddress	`json:"additionalDnAaaAddresses,omitempty"`
	 ThreeGppChargingCharacteristics	string	`json:"3gppChargingCharacteristics,omitempty"`
	 PduSessionContinuityInd	PduSessionContinuityInd	`json:"pduSessionContinuityInd,omitempty"`
	 AcsInfo	*AcsInfo	`json:"acsInfo,omitempty"`
	 SecondaryAuth	*bool	`json:"secondaryAuth,omitempty"`
	 DnAaaIpAddressAllocation	*bool	`json:"dnAaaIpAddressAllocation,omitempty"`
}
