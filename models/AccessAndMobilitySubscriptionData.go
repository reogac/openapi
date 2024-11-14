/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
}
