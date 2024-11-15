/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateData struct {
	RequestIndication                    RequestIndication             `json:"requestIndication"`
	IcnTunnelInfo                        *TunnelInfo                   `json:"icnTunnelInfo,omitempty"`
	UnknownN1SmInfo                      *RefToBinaryData              `json:"unknownN1SmInfo,omitempty"`
	AlwaysOnRequested                    *bool                         `json:"alwaysOnRequested,omitempty"`
	MaRequestInd                         *bool                         `json:"maRequestInd,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo            `json:"pcfUeCallbackInfo,omitempty"`
	DisasterRoamingInd                   *DisasterRoamingInd           `json:"disasterRoamingInd,omitempty"`
	SecondaryRatUsageDataReportContainer []string                      `json:"secondaryRatUsageDataReportContainer,omitempty"`
	ServingNetwork                       *PlmnIdNid                    `json:"servingNetwork,omitempty"`
	AnType                               AccessType                    `json:"anType,omitempty"`
	PauseCharging                        *bool                         `json:"pauseCharging,omitempty"`
	FiveGMmCauseValue                    *int                          `json:"5gMmCauseValue,omitempty"`
	RoamingChargingProfile               *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter             `json:"moExpDataCounter,omitempty"`
	Guami                                *Guami                        `json:"guami,omitempty"`
	MaxIntegrityProtectedDataRateDl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	VcnTunnelInfo                        *TunnelInfo                   `json:"vcnTunnelInfo,omitempty"`
	AnTypeCanBeChanged                   *bool                         `json:"anTypeCanBeChanged,omitempty"`
	PsaInfo                              []PsaInformation              `json:"psaInfo,omitempty"`
	IsmfId                               string                        `json:"ismfId,omitempty"`
	MaxIntegrityProtectedDataRateUl      MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	AdditionalCnTunnelInfo               *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	QosFlowsNotifyList                   []QosFlowNotifyItem           `json:"qosFlowsNotifyList,omitempty"`
	RevokeEbiList                        []int                         `json:"revokeEbiList,omitempty"`
	Cause                                Cause                         `json:"cause,omitempty"`
	NotifyList                           []PduSessionNotifyItem        `json:"NotifyList,omitempty"`
	UnavailableAccessInd                 UnavailableAccessIndication   `json:"unavailableAccessInd,omitempty"`
	IsmfPduSessionUri                    string                        `json:"ismfPduSessionUri,omitempty"`
	SupportedFeatures                    string                        `json:"supportedFeatures,omitempty"`
	DlServingPlmnRateCtl                 *int                          `json:"dlServingPlmnRateCtl,omitempty"`
	UeLocation                           *UserLocation                 `json:"ueLocation,omitempty"`
	N1SmInfoFromUe                       *RefToBinaryData              `json:"n1SmInfoFromUe,omitempty"`
	VsmfId                               string                        `json:"vsmfId,omitempty"`
	ISmfServiceInstanceId                string                        `json:"iSmfServiceInstanceId,omitempty"`
	VplmnQos                             *VplmnQos                     `json:"vplmnQos,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd            `json:"smPolicyNotifyInd,omitempty"`
	HoPreparationIndication              *bool                         `json:"hoPreparationIndication,omitempty"`
	SecurityResult                       *SecurityResult               `json:"securityResult,omitempty"`
	UeTimeZone                           string                        `json:"ueTimeZone,omitempty"`
	UpCnxState                           UpCnxState                    `json:"upCnxState,omitempty"`
	DnaiList                             []string                      `json:"dnaiList,omitempty"`
	AmfNfId                              string                        `json:"amfNfId,omitempty"`
	EpsBearerId                          []int                         `json:"epsBearerId,omitempty"`
	NgApCause                            *NgApCause                    `json:"ngApCause,omitempty"`
	SecondaryRatUsageInfo                []SecondaryRatUsageInfo       `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt1                           *N4Information                `json:"n4InfoExt1,omitempty"`
	N4InfoExt2                           *N4Information                `json:"n4InfoExt2,omitempty"`
	VsmfPduSessionUri                    string                        `json:"vsmfPduSessionUri,omitempty"`
	AddUeLocation                        *UserLocation                 `json:"addUeLocation,omitempty"`
	MaReleaseInd                         MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	MaNwUpgradeInd                       *bool                         `json:"maNwUpgradeInd,omitempty"`
	UpSecurityInfo                       *UpSecurityInfo               `json:"upSecurityInfo,omitempty"`
	Pei                                  string                        `json:"pei,omitempty"`
	N4Info                               *N4Information                `json:"n4Info,omitempty"`
	PresenceInLadn                       PresenceState                 `json:"presenceInLadn,omitempty"`
	VSmfServiceInstanceId                string                        `json:"vSmfServiceInstanceId,omitempty"`
	RatType                              RatType                       `json:"ratType,omitempty"`
	QosFlowsRelNotifyList                []QosFlowItem                 `json:"qosFlowsRelNotifyList,omitempty"`
	SecondaryRatUsageReport              []SecondaryRatUsageReport     `json:"secondaryRatUsageReport,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory     `json:"satelliteBackhaulCat,omitempty"`
	AdditionalAnType                     AccessType                    `json:"additionalAnType,omitempty"`
	Pti                                  *int                          `json:"pti,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication     `json:"epsInterworkingInd,omitempty"`
	UlclBpInfo                           *UlclBpInformation            `json:"ulclBpInfo,omitempty"`
}
