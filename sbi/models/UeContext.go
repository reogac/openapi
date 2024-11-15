/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContext struct {
	PcfGroupId                 string                          `json:"pcfGroupId,omitempty"`
	PcfUeAmbr                  *Ambr                           `json:"pcfUeAmbr,omitempty"`
	FiveGMmCapability          string                          `json:"5gMmCapability,omitempty"`
	CMsisdn                    string                          `json:"cMsisdn,omitempty"`
	V2xContext                 *V2xContext                     `json:"v2xContext,omitempty"`
	SubRfsp                    *int                            `json:"subRfsp,omitempty"`
	SessionContextList         []PduSessionContext             `json:"sessionContextList,omitempty"`
	EcRestrictionDataNb        *bool                           `json:"ecRestrictionDataNb,omitempty"`
	ProseContext               *ProseContext                   `json:"proseContext,omitempty"`
	PcfRfsp                    *int                            `json:"pcfRfsp,omitempty"`
	SubUeAmbr                  *Ambr                           `json:"subUeAmbr,omitempty"`
	UePolicyReqTriggerList     []string                        `json:"uePolicyReqTriggerList,omitempty"`
	MoExpDataCounter           *MoExpDataCounter               `json:"moExpDataCounter,omitempty"`
	PcfUepBindingInfo          string                          `json:"pcfUepBindingInfo,omitempty"`
	MsClassmark2               string                          `json:"msClassmark2,omitempty"`
	SeafData                   *SeafData                       `json:"seafData,omitempty"`
	PcfSetId                   string                          `json:"pcfSetId,omitempty"`
	PcfBinding                 SbiBindingLevel                 `json:"pcfBinding,omitempty"`
	PcfUePolicyUri             string                          `json:"pcfUePolicyUri,omitempty"`
	ServiceAreaRestriction     *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	SnpnOnboardInd             *bool                           `json:"snpnOnboardInd,omitempty"`
	UdmGroupId                 string                          `json:"udmGroupId,omitempty"`
	PcfId                      string                          `json:"pcfId,omitempty"`
	PcfAmpServiceSetId         string                          `json:"pcfAmpServiceSetId,omitempty"`
	UpdpSubscriptionData       *UpdpSubscriptionData           `json:"updpSubscriptionData,omitempty"`
	AstiDistributionIndication *bool                           `json:"astiDistributionIndication,omitempty"`
	Supi                       string                          `json:"supi,omitempty"`
	AnalyticsSubscriptionList  []AnalyticsSubscription         `json:"analyticsSubscriptionList,omitempty"`
	PraInAmPolicy              map[string]PresenceInfo         `json:"praInAmPolicy,omitempty"`
	SupiUnauthInd              *bool                           `json:"supiUnauthInd,omitempty"`
	Pei                        string                          `json:"pei,omitempty"`
	HNwPubKeyId                *int                            `json:"hNwPubKeyId,omitempty"`
	PcfUeSliceMbrList          map[string]SliceMbr             `json:"pcfUeSliceMbrList,omitempty"`
	HpcfId                     string                          `json:"hpcfId,omitempty"`
	ImmediateMdtConf           *ImmediateMdtConf               `json:"immediateMdtConf,omitempty"`
	ManagementMdtInd           *bool                           `json:"managementMdtInd,omitempty"`
	UsedServiceAreaRestriction *ServiceAreaRestriction         `json:"usedServiceAreaRestriction,omitempty"`
	SmsfSetId                  string                          `json:"smsfSetId,omitempty"`
	GroupList                  []string                        `json:"groupList,omitempty"`
	DrxParameter               string                          `json:"drxParameter,omitempty"`
	SubUeSliceMbrList          map[string]SliceMbr             `json:"subUeSliceMbrList,omitempty"`
	AmPolicyReqTriggerList     []string                        `json:"amPolicyReqTriggerList,omitempty"`
	ForbiddenAreaList          []Area                          `json:"forbiddenAreaList,omitempty"`
	SmsfServiceSetId           string                          `json:"smsfServiceSetId,omitempty"`
	IabOperationAllowed        *bool                           `json:"iabOperationAllowed,omitempty"`
	PraInUePolicy              map[string]PresenceInfo         `json:"praInUePolicy,omitempty"`
	TsErrorBudget              *int                            `json:"tsErrorBudget,omitempty"`
	GpsiList                   []string                        `json:"gpsiList,omitempty"`
	EpsInterworkingInfo        *EpsInterworkingInfo            `json:"epsInterworkingInfo,omitempty"`
	ServiceGapExpiryTime       string                          `json:"serviceGapExpiryTime,omitempty"`
	StnSr                      string                          `json:"stnSr,omitempty"`
	SmallDataRateStatusInfos   []SmallDataRateStatusInfo       `json:"smallDataRateStatusInfos,omitempty"`
	DisasterRoamingInd         *bool                           `json:"disasterRoamingInd,omitempty"`
	DisasterPlmn               *PlmnId                         `json:"disasterPlmn,omitempty"`
	SmfSelInfo                 *SmfSelectionData               `json:"smfSelInfo,omitempty"`
	PcfAmPolicyUri             string                          `json:"pcfAmPolicyUri,omitempty"`
	CagData                    *CagData                        `json:"cagData,omitempty"`
	EcRestrictionDataWb        *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	PcfAmpBindingInfo          string                          `json:"pcfAmpBindingInfo,omitempty"`
	SmPolicyNotifyPduList      []PduSessionInfo                `json:"smPolicyNotifyPduList,omitempty"`
	UsedRfsp                   *int                            `json:"usedRfsp,omitempty"`
	EventSubscriptionList      []ExtAmfEventSubscription       `json:"eventSubscriptionList,omitempty"`
	MmContextList              []MmContext                     `json:"mmContextList,omitempty"`
	PcfUeCallbackInfo          *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	WlServAreaRes              *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	SmsfBindingInfo            string                          `json:"smsfBindingInfo,omitempty"`
	AsTimeDisParam             *AsTimeDistributionParam        `json:"asTimeDisParam,omitempty"`
	RoutingIndicator           string                          `json:"routingIndicator,omitempty"`
	SmsfId                     string                          `json:"smsfId,omitempty"`
	PcfUepServiceSetId         string                          `json:"pcfUepServiceSetId,omitempty"`
	RestrictedCoreNwTypeList   []string                        `json:"restrictedCoreNwTypeList,omitempty"`
	SupportedCodecList         []string                        `json:"supportedCodecList,omitempty"`
	AusfGroupId                string                          `json:"ausfGroupId,omitempty"`
	HpcfSetId                  string                          `json:"hpcfSetId,omitempty"`
	RedCapInd                  *bool                           `json:"redCapInd,omitempty"`
	RestrictedRatList          []string                        `json:"restrictedRatList,omitempty"`
	TraceData                  *TraceData                      `json:"traceData,omitempty"`
	RestrictedPrimaryRatList   []string                        `json:"restrictedPrimaryRatList,omitempty"`
	RestrictedSecondaryRatList []string                        `json:"restrictedSecondaryRatList,omitempty"`
	LteCatMInd                 *bool                           `json:"lteCatMInd,omitempty"`
	UePositioningCap           string                          `json:"uePositioningCap,omitempty"`
}
