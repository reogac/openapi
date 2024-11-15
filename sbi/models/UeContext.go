/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
}
