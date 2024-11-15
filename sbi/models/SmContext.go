/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	PcfId                           string                          `json:"pcfId,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
}
