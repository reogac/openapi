/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
}
