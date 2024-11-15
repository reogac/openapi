/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
	Dnn                             string                          `json:"dnn"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
}
