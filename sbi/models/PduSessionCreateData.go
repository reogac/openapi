/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateData struct {
	UnauthenticatedSupi             *bool                           `json:"unauthenticatedSupi,omitempty"`
	Dnn                             string                          `json:"dnn"`
	IsmfId                          string                          `json:"ismfId,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AmfNfId                         string                          `json:"amfNfId,omitempty"`
	Pei                             string                          `json:"pei,omitempty"`
	AnType                          AccessType                      `json:"anType"`
	AdditionalAnType                AccessType                      `json:"additionalAnType,omitempty"`
	AddUeLocation                   *UserLocation                   `json:"addUeLocation,omitempty"`
	HPcfId                          string                          `json:"hPcfId,omitempty"`
	HoPreparationIndication         *bool                           `json:"hoPreparationIndication,omitempty"`
	ISmfServiceInstanceId           string                          `json:"iSmfServiceInstanceId,omitempty"`
	InvokeNef                       *bool                           `json:"invokeNef,omitempty"`
	MaNwUpgradeInd                  *bool                           `json:"maNwUpgradeInd,omitempty"`
	SecondaryRatUsageInfo           []SecondaryRatUsageInfo         `json:"secondaryRatUsageInfo,omitempty"`
	OldPduSessionRef                string                          `json:"oldPduSessionRef,omitempty"`
	UpipSupported                   *bool                           `json:"upipSupported,omitempty"`
	N9ForwardingTunnelInfo          *TunnelInfo                     `json:"n9ForwardingTunnelInfo,omitempty"`
	RatType                         RatType                         `json:"ratType,omitempty"`
	CpCiotEnabled                   *bool                           `json:"cpCiotEnabled,omitempty"`
	ApnRateStatus                   *ApnRateStatus                  `json:"apnRateStatus,omitempty"`
	IsmfPduSessionUri               string                          `json:"ismfPduSessionUri,omitempty"`
	UeTimeZone                      string                          `json:"ueTimeZone,omitempty"`
	N1SmInfoFromUe                  *RefToBinaryData                `json:"n1SmInfoFromUe,omitempty"`
	EpsBearerCtxStatus              string                          `json:"epsBearerCtxStatus,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	OldSmContextRef                 string                          `json:"oldSmContextRef,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	VcnTunnelInfo                   *TunnelInfo                     `json:"vcnTunnelInfo,omitempty"`
	UnknownN1SmInfo                 *RefToBinaryData                `json:"unknownN1SmInfo,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	PresenceInLadn                  PresenceState                   `json:"presenceInLadn,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	CpOnlyInd                       *bool                           `json:"cpOnlyInd,omitempty"`
	PcfUeCallbackInfo               *PcfUeCallbackInfo              `json:"pcfUeCallbackInfo,omitempty"`
	VSmfServiceInstanceId           string                          `json:"vSmfServiceInstanceId,omitempty"`
	DlServingPlmnRateCtl            *int                            `json:"dlServingPlmnRateCtl,omitempty"`
	UpSecurityInfo                  *UpSecurityInfo                 `json:"upSecurityInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	VplmnQos                        *VplmnQos                       `json:"vplmnQos,omitempty"`
	VsmfPduSessionUri               string                          `json:"vsmfPduSessionUri,omitempty"`
	SmallDataRateStatus             *SmallDataRateStatus            `json:"smallDataRateStatus,omitempty"`
	VsmfId                          string                          `json:"vsmfId,omitempty"`
	PgwS8cFteid                     string                          `json:"pgwS8cFteid,omitempty"`
	EpsInterworkingInd              EpsInterworkingIndication       `json:"epsInterworkingInd,omitempty"`
	DisasterRoamingInd              *bool                           `json:"disasterRoamingInd,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	EpsBearerId                     []int                           `json:"epsBearerId,omitempty"`
	IcnTunnelInfo                   *TunnelInfo                     `json:"icnTunnelInfo,omitempty"`
	UeLocation                      *UserLocation                   `json:"ueLocation,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Supi                            string                          `json:"supi,omitempty"`
	ServingNetwork                  PlmnIdNid                       `json:"servingNetwork"`
	RequestType                     RequestType                     `json:"requestType,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SmPolicyNotifyInd               *bool                           `json:"smPolicyNotifyInd,omitempty"`
	UpCnxState                      UpCnxState                      `json:"upCnxState,omitempty"`
	AlwaysOnRequested               *bool                           `json:"alwaysOnRequested,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	OldPduSessionId                 *int                            `json:"oldPduSessionId,omitempty"`
	Guami                           *Guami                          `json:"guami,omitempty"`
	MaRequestInd                    *bool                           `json:"maRequestInd,omitempty"`
}
