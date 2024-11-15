/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContext struct {
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	SNssai                          Snssai                          `json:"sNssai"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	SatelliteBackhaulCat            SatelliteBackhaulCategory       `json:"satelliteBackhaulCat,omitempty"`
	SelMode                         DnnSelectionMode                `json:"selMode,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	PduSessionSmfBinding            SbiBindingLevel                 `json:"pduSessionSmfBinding,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
}
