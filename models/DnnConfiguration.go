package models

type DnnConfiguration struct {
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
}
