/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
}
