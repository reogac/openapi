/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	CagData                         *CagData                        `json:"cagData,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
}
