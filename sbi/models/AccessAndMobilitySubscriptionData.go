/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccessAndMobilitySubscriptionData struct {
	CagData                         *CagData                        `json:"cagData,omitempty"`
	SubsRegTimer                    *int                            `json:"subsRegTimer,omitempty"`
	SorafRetrieval                  *bool                           `json:"sorafRetrieval,omitempty"`
	EcRestrictionDataWb             *EcRestrictionDataWb            `json:"ecRestrictionDataWb,omitempty"`
	HssGroupId                      string                          `json:"hssGroupId,omitempty"`
	NbIoTUePriority                 *int                            `json:"nbIoTUePriority,omitempty"`
	EdrxParametersList              []EdrxParameters                `json:"edrxParametersList,omitempty"`
	ThreeGppChargingCharacteristics string                          `json:"3gppChargingCharacteristics,omitempty"`
	InternalGroupIds                []string                        `json:"internalGroupIds,omitempty"`
	MicoAllowed                     *bool                           `json:"micoAllowed,omitempty"`
	MdtUserConsent                  MdtUserConsent                  `json:"mdtUserConsent,omitempty"`
	AerialUeSubInfo                 *AerialUeSubscriptionInfo       `json:"aerialUeSubInfo,omitempty"`
	ExpectedUeBehaviourList         *ExpectedUeBehaviourData        `json:"expectedUeBehaviourList,omitempty"`
	SecondaryRatRestrictions        []string                        `json:"secondaryRatRestrictions,omitempty"`
	UpuInfo                         *UpuInfo                        `json:"upuInfo,omitempty"`
	SubscribedDnnList               []string                        `json:"subscribedDnnList,omitempty"`
	ServiceGapTime                  *int                            `json:"serviceGapTime,omitempty"`
	AdjacentPlmnRestrictions        map[string]PlmnRestriction      `json:"adjacentPlmnRestrictions,omitempty"`
	Nssai                           *Nssai                          `json:"nssai,omitempty"`
	ServiceAreaRestriction          *ServiceAreaRestriction         `json:"serviceAreaRestriction,omitempty"`
	ActiveTime                      *int                            `json:"activeTime,omitempty"`
	OdbPacketServices               OdbPacketServices               `json:"odbPacketServices,omitempty"`
	NssaiInclusionAllowed           *bool                           `json:"nssaiInclusionAllowed,omitempty"`
	IabOperationAllowed             *bool                           `json:"iabOperationAllowed,omitempty"`
	EcRestrictionDataNb             *bool                           `json:"ecRestrictionDataNb,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	Gpsis                           []string                        `json:"gpsis,omitempty"`
	CoreNetworkTypeRestrictions     []string                        `json:"coreNetworkTypeRestrictions,omitempty"`
	RfspIndex                       *int                            `json:"rfspIndex,omitempty"`
	SharedAmDataIds                 []string                        `json:"sharedAmDataIds,omitempty"`
	SharedVnGroupDataIds            map[string]string               `json:"sharedVnGroupDataIds,omitempty"`
	SorInfoExpectInd                *bool                           `json:"sorInfoExpectInd,omitempty"`
	RatRestrictions                 []string                        `json:"ratRestrictions,omitempty"`
	UeUsageType                     *int                            `json:"ueUsageType,omitempty"`
	SorInfo                         *SorInfo                        `json:"sorInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	CMsisdn                         string                          `json:"cMsisdn,omitempty"`
	MpsPriority                     *bool                           `json:"mpsPriority,omitempty"`
	TraceData                       *TraceData                      `json:"traceData,omitempty"`
	StnSr                           string                          `json:"stnSr,omitempty"`
	WirelineForbiddenAreas          []WirelineArea                  `json:"wirelineForbiddenAreas,omitempty"`
	RemoteProvInd                   *bool                           `json:"remoteProvInd,omitempty"`
	SubscribedUeAmbr                *Ambr                           `json:"subscribedUeAmbr,omitempty"`
	SorUpdateIndicatorList          []string                        `json:"sorUpdateIndicatorList,omitempty"`
	RgWirelineCharacteristics       string                          `json:"rgWirelineCharacteristics,omitempty"`
	PrimaryRatRestrictions          []string                        `json:"primaryRatRestrictions,omitempty"`
	PtwParametersList               []PtwParameters                 `json:"ptwParametersList,omitempty"`
	McsPriority                     *bool                           `json:"mcsPriority,omitempty"`
	MdtConfiguration                *MdtConfiguration               `json:"mdtConfiguration,omitempty"`
	WirelineServiceAreaRestriction  *WirelineServiceAreaRestriction `json:"wirelineServiceAreaRestriction,omitempty"`
	PcfSelectionAssistanceInfos     []PcfSelectionAssistanceInfo    `json:"pcfSelectionAssistanceInfos,omitempty"`
	ForbiddenAreas                  []Area                          `json:"forbiddenAreas,omitempty"`
	RoamingRestrictions             *RoamingRestrictions            `json:"roamingRestrictions,omitempty"`
}
