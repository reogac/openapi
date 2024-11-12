package models

type DnnConfiguration struct {
	SscModes                             SscModes                `json:"sscModes"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
}
