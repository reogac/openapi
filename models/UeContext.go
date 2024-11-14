/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
}
