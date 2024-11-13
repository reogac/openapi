package models

type DnnConfiguration struct {
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
}
