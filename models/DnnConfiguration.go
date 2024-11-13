package models

type DnnConfiguration struct {
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
}
