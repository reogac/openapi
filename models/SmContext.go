/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	SmfUri                          string                          `json:"smfUri,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	Dnn                             string                          `json:"dnn"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
}
