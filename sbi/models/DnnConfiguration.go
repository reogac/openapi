/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnnConfiguration struct {
	EasDiscoveryAuthorized               *bool                   `json:"easDiscoveryAuthorized,omitempty"`
	SscModes                             SscModes                `json:"sscModes"`
	SessionAmbr                          *Ambr                   `json:"sessionAmbr,omitempty"`
	NiddNefId                            string                  `json:"niddNefId,omitempty"`
	RedundantSessionAllowed              *bool                   `json:"redundantSessionAllowed,omitempty"`
	AtsssAllowed                         *bool                   `json:"atsssAllowed,omitempty"`
	AerialUeInd                          AerialUeIndication      `json:"aerialUeInd,omitempty"`
	IwkEpsInd                            *bool                   `json:"iwkEpsInd,omitempty"`
	UpSecurity                           *UpSecurity             `json:"upSecurity,omitempty"`
	UavSecondaryAuth                     *bool                   `json:"uavSecondaryAuth,omitempty"`
	DnAaaAddress                         *IpAddress              `json:"dnAaaAddress,omitempty"`
	SharedEcsAddrConfigInfo              string                  `json:"sharedEcsAddrConfigInfo,omitempty"`
	FiveGQosProfile                      *SubscribedDefaultQos   `json:"5gQosProfile,omitempty"`
	DnAaaIpAddressAllocation             *bool                   `json:"dnAaaIpAddressAllocation,omitempty"`
	Ipv6Index                            *IpIndex                `json:"ipv6Index,omitempty"`
	OnboardingInd                        *bool                   `json:"onboardingInd,omitempty"`
	ThreeGppChargingCharacteristics      string                  `json:"3gppChargingCharacteristics,omitempty"`
	StaticIpAddress                      []IpAddress             `json:"staticIpAddress,omitempty"`
	NiddInfo                             *NiddInformation        `json:"niddInfo,omitempty"`
	Ipv4FrameRouteList                   []FrameRouteInfo        `json:"ipv4FrameRouteList,omitempty"`
	Ipv4Index                            *IpIndex                `json:"ipv4Index,omitempty"`
	PduSessionContinuityInd              PduSessionContinuityInd `json:"pduSessionContinuityInd,omitempty"`
	Ipv6FrameRouteList                   []FrameRouteInfo        `json:"ipv6FrameRouteList,omitempty"`
	AdditionalDnAaaAddresses             []IpAddress             `json:"additionalDnAaaAddresses,omitempty"`
	AdditionalEcsAddrConfigInfos         []EcsAddrConfigInfo     `json:"additionalEcsAddrConfigInfos,omitempty"`
	PduSessionTypes                      PduSessionTypes         `json:"pduSessionTypes"`
	AdditionalSharedEcsAddrConfigInfoIds []string                `json:"additionalSharedEcsAddrConfigInfoIds,omitempty"`
	AcsInfo                              *AcsInfo                `json:"acsInfo,omitempty"`
	SecondaryAuth                        *bool                   `json:"secondaryAuth,omitempty"`
	IptvAccCtrlInfo                      string                  `json:"iptvAccCtrlInfo,omitempty"`
	EcsAddrConfigInfo                    *EcsAddrConfigInfo      `json:"ecsAddrConfigInfo,omitempty"`
	SubscribedMaxIpv6PrefixSize          *int                    `json:"subscribedMaxIpv6PrefixSize,omitempty"`
	DnAaaFqdn                            string                  `json:"dnAaaFqdn,omitempty"`
}
