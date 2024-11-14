/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
}
