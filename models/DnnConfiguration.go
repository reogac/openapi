package models

type DnnConfiguration struct {
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
}
