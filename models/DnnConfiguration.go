package models

type DnnConfiguration struct {
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
}
