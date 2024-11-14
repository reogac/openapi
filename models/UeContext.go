/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
}
