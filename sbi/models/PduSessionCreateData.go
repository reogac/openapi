/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	Dnn                             string                          `json:"dnn"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
}
