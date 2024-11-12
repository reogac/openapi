package models

type DnnConfiguration struct {
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
}
