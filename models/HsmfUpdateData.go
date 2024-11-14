/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateData struct {
	AdditionalCnTunnelInfo               *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	SecondaryRatUsageInfo                []SecondaryRatUsageInfo       `json:"secondaryRatUsageInfo,omitempty"`
	PsaInfo                              []PsaInformation              `json:"psaInfo,omitempty"`
	VsmfId                               string                        `json:"vsmfId,omitempty"`
	MaxIntegrityProtectedDataRateUl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	Pei                                  string                        `json:"pei,omitempty"`
	AdditionalAnType                     AccessType                    `json:"additionalAnType,omitempty"`
	MaReleaseInd                         MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	N4InfoExt1                           *N4Information                `json:"n4InfoExt1,omitempty"`
	SecondaryRatUsageReport              []SecondaryRatUsageReport     `json:"secondaryRatUsageReport,omitempty"`
	N4Info                               *N4Information                `json:"n4Info,omitempty"`
	ISmfServiceInstanceId                string                        `json:"iSmfServiceInstanceId,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd            `json:"smPolicyNotifyInd,omitempty"`
	VcnTunnelInfo                        *TunnelInfo                   `json:"vcnTunnelInfo,omitempty"`
	AnType                               AccessType                    `json:"anType,omitempty"`
	EpsBearerId                          []int                         `json:"epsBearerId,omitempty"`
	Cause                                Cause                         `json:"cause,omitempty"`
	UeTimeZone                           string                        `json:"ueTimeZone,omitempty"`
	AddUeLocation                        *UserLocation                 `json:"addUeLocation,omitempty"`
	PauseCharging                        *bool                         `json:"pauseCharging,omitempty"`
	UpSecurityInfo                       *UpSecurityInfo               `json:"upSecurityInfo,omitempty"`
	RevokeEbiList                        []int                         `json:"revokeEbiList,omitempty"`
	RatType                              RatType                       `json:"ratType,omitempty"`
	HoPreparationIndication              *bool                         `json:"hoPreparationIndication,omitempty"`
	VsmfPduSessionUri                    string                        `json:"vsmfPduSessionUri,omitempty"`
	UeLocation                           *UserLocation                 `json:"ueLocation,omitempty"`
	UpCnxState                           UpCnxState                    `json:"upCnxState,omitempty"`
	DlServingPlmnRateCtl                 *int                          `json:"dlServingPlmnRateCtl,omitempty"`
	DnaiList                             []string                      `json:"dnaiList,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory     `json:"satelliteBackhaulCat,omitempty"`
	AnTypeCanBeChanged                   *bool                         `json:"anTypeCanBeChanged,omitempty"`
	MaRequestInd                         *bool                         `json:"maRequestInd,omitempty"`
	UnavailableAccessInd                 UnavailableAccessIndication   `json:"unavailableAccessInd,omitempty"`
	IsmfId                               string                        `json:"ismfId,omitempty"`
	N1SmInfoFromUe                       *RefToBinaryData              `json:"n1SmInfoFromUe,omitempty"`
	NgApCause                            *NgApCause                    `json:"ngApCause,omitempty"`
	NotifyList                           []PduSessionNotifyItem        `json:"NotifyList,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter             `json:"moExpDataCounter,omitempty"`
	Guami                                *Guami                        `json:"guami,omitempty"`
	DisasterRoamingInd                   *DisasterRoamingInd           `json:"disasterRoamingInd,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication     `json:"epsInterworkingInd,omitempty"`
	N4InfoExt2                           *N4Information                `json:"n4InfoExt2,omitempty"`
	AmfNfId                              string                        `json:"amfNfId,omitempty"`
	UnknownN1SmInfo                      *RefToBinaryData              `json:"unknownN1SmInfo,omitempty"`
	QosFlowsRelNotifyList                []QosFlowItem                 `json:"qosFlowsRelNotifyList,omitempty"`
	FiveGMmCauseValue                    *int                          `json:"5gMmCauseValue,omitempty"`
	AlwaysOnRequested                    *bool                         `json:"alwaysOnRequested,omitempty"`
	VSmfServiceInstanceId                string                        `json:"vSmfServiceInstanceId,omitempty"`
	IsmfPduSessionUri                    string                        `json:"ismfPduSessionUri,omitempty"`
	IcnTunnelInfo                        *TunnelInfo                   `json:"icnTunnelInfo,omitempty"`
	ServingNetwork                       *PlmnIdNid                    `json:"servingNetwork,omitempty"`
	MaNwUpgradeInd                       *bool                         `json:"maNwUpgradeInd,omitempty"`
	PresenceInLadn                       PresenceState                 `json:"presenceInLadn,omitempty"`
	MaxIntegrityProtectedDataRateDl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RequestIndication                    RequestIndication             `json:"requestIndication"`
	Pti                                  *int                          `json:"pti,omitempty"`
	SupportedFeatures                    string                        `json:"supportedFeatures,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo            `json:"pcfUeCallbackInfo,omitempty"`
	VplmnQos                             *VplmnQos                     `json:"vplmnQos,omitempty"`
	SecurityResult                       *SecurityResult               `json:"securityResult,omitempty"`
	SecondaryRatUsageDataReportContainer []string                      `json:"secondaryRatUsageDataReportContainer,omitempty"`
	QosFlowsNotifyList                   []QosFlowNotifyItem           `json:"qosFlowsNotifyList,omitempty"`
	UlclBpInfo                           *UlclBpInformation            `json:"ulclBpInfo,omitempty"`
	RoamingChargingProfile               *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
}
