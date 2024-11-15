/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdateData struct {
	SatelliteBackhaulCat                 SatelliteBackhaulCategory          `json:"satelliteBackhaulCat,omitempty"`
	ServingNfId                          string                             `json:"servingNfId,omitempty"`
	AnType                               AccessType                         `json:"anType,omitempty"`
	N2SmInfo                             *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N2SmInfoType                         N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	MaNwUpgradeInd                       *bool                              `json:"maNwUpgradeInd,omitempty"`
	Guami                                *Guami                             `json:"guami,omitempty"`
	ServingNetwork                       *PlmnIdNid                         `json:"servingNetwork,omitempty"`
	DataForwarding                       *bool                              `json:"dataForwarding,omitempty"`
	MaRequestInd                         *bool                              `json:"maRequestInd,omitempty"`
	ForwardingFTeid                      string                             `json:"forwardingFTeid,omitempty"`
	ToBeSwitched                         *bool                              `json:"toBeSwitched,omitempty"`
	N9DlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9DlForwardingTnlList,omitempty"`
	FiveGMmCauseValue                    *int                               `json:"5gMmCauseValue,omitempty"`
	EpsBearerSetup                       []string                           `json:"epsBearerSetup,omitempty"`
	N2SmInfoExt1                         *RefToBinaryData                   `json:"n2SmInfoExt1,omitempty"`
	SupportedFeatures                    string                             `json:"supportedFeatures,omitempty"`
	AnTypeToReactivate                   AccessType                         `json:"anTypeToReactivate,omitempty"`
	UeTimeZone                           string                             `json:"ueTimeZone,omitempty"`
	MaReleaseInd                         MaReleaseIndication                `json:"maReleaseInd,omitempty"`
	ExtendedNasSmTimerInd                *bool                              `json:"extendedNasSmTimerInd,omitempty"`
	AddUeLocation                        *UserLocation                      `json:"addUeLocation,omitempty"`
	AnTypeCanBeChanged                   *bool                              `json:"anTypeCanBeChanged,omitempty"`
	N2SmInfoTypeExt1                     N2SmInfoType                       `json:"n2SmInfoTypeExt1,omitempty"`
	AdditionalAnType                     AccessType                         `json:"additionalAnType,omitempty"`
	UpCnxState                           UpCnxState                         `json:"upCnxState,omitempty"`
	EpsInterworkingInd                   EpsInterworkingIndication          `json:"epsInterworkingInd,omitempty"`
	MoExpDataCounter                     *MoExpDataCounter                  `json:"moExpDataCounter,omitempty"`
	N1SmMsg                              *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	ExemptionInd                         *ExemptionInd                      `json:"exemptionInd,omitempty"`
	DdnFailureSubs                       *DdnFailureSubs                    `json:"ddnFailureSubs,omitempty"`
	Pei                                  string                             `json:"pei,omitempty"`
	PresenceInLadn                       PresenceState                      `json:"presenceInLadn,omitempty"`
	FailedToBeSwitched                   *bool                              `json:"failedToBeSwitched,omitempty"`
	N9InactivityTimer                    *int                               `json:"n9InactivityTimer,omitempty"`
	RevokeEbiList                        []int                              `json:"revokeEbiList,omitempty"`
	BackupAmfInfo                        []BackupAmfInfo                    `json:"backupAmfInfo,omitempty"`
	N9ForwardingTunnel                   *TunnelInfo                        `json:"n9ForwardingTunnel,omitempty"`
	N9UlForwardingTnlList                []IndirectDataForwardingTunnelInfo `json:"n9UlForwardingTnlList,omitempty"`
	NgApCause                            *NgApCause                         `json:"ngApCause,omitempty"`
	SNssai                               *Snssai                            `json:"sNssai,omitempty"`
	RatType                              RatType                            `json:"ratType,omitempty"`
	Release                              *bool                              `json:"release,omitempty"`
	SecondaryRatUsageDataReportContainer []string                           `json:"secondaryRatUsageDataReportContainer,omitempty"`
	PcfUeCallbackInfo                    *PcfUeCallbackInfo                 `json:"pcfUeCallbackInfo,omitempty"`
	UeLocation                           *UserLocation                      `json:"ueLocation,omitempty"`
	ForwardingBearerContexts             []string                           `json:"forwardingBearerContexts,omitempty"`
	HoState                              HoState                            `json:"hoState,omitempty"`
	TraceData                            *TraceData                         `json:"traceData,omitempty"`
	SkipN2PduSessionResRelInd            *bool                              `json:"skipN2PduSessionResRelInd,omitempty"`
	SmPolicyNotifyInd                    *SmPolicyNotifyInd                 `json:"smPolicyNotifyInd,omitempty"`
	TargetId                             *NgRanTargetId                     `json:"targetId,omitempty"`
	TargetServingNfId                    string                             `json:"targetServingNfId,omitempty"`
	SmContextStatusUri                   string                             `json:"smContextStatusUri,omitempty"`
	N9DlForwardingTunnel                 *TunnelInfo                        `json:"n9DlForwardingTunnel,omitempty"`
	Cause                                Cause                              `json:"cause,omitempty"`
}
