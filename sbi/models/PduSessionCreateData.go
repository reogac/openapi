/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
}
