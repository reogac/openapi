/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	Guami                                *Guami                             `json:"guami,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	AnType                               AccessType                         `json:"anType,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
}
