package models

type SmContext struct {
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	SelectedDnn                     string                          `json:"selectedDnn,omitempty"`
	SmfUri                          string                          `json:"smfUri,omitempty"`
	QosFlowsList                    []QosFlowSetupItem              `json:"qosFlowsList"`
	PduSessionSmfServiceSetId       string                          `json:"pduSessionSmfServiceSetId,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	HNwPubKeyId                     *int                            `json:"hNwPubKeyId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	ChargingId                      string                          `json:"chargingId,omitempty"`
	ChargingInfo                    *ChargingInformation            `json:"chargingInfo,omitempty"`
	AddRanTunnelInfo                []QosFlowTunnel                 `json:"addRanTunnelInfo,omitempty"`
	RedRanTunnelInfo                *QosFlowTunnel                  `json:"redRanTunnelInfo,omitempty"`
	AnchorSmfOauth2Required         *bool                           `json:"anchorSmfOauth2Required,omitempty"`
	PcfGroupId                      string                          `json:"pcfGroupId,omitempty"`
	MaxIntegrityProtectedDataRate   string                          `json:"maxIntegrityProtectedDataRate,omitempty"`
	N9fscSupportInd                 *bool                           `json:"n9fscSupportInd,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	DlsetSupportInd                 *bool                           `json:"dlsetSupportInd,omitempty"`
	RoutingIndicator                string                          `json:"routingIndicator,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	PcfSetId                        string                          `json:"pcfSetId,omitempty"`
	PduSessionSmfSetId              string                          `json:"pduSessionSmfSetId,omitempty"`
	Dnn                             string                          `json:"dnn"`
	HSmfUri                         string                          `json:"hSmfUri,omitempty"`
	RanTunnelInfo                   *QosFlowTunnel                  `json:"ranTunnelInfo,omitempty"`
	SatelliteBackhaulCat            string                          `json:"satelliteBackhaulCat,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	PcfId                           string                          `json:"pcfId,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	ForwardingInd                   *bool                           `json:"forwardingInd,omitempty"`
	HplmnSnssai                     *Snssai                         `json:"hplmnSnssai,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	SscMode                         string                          `json:"sscMode,omitempty"`
	SessionAmbr                     Ambr                            `json:"sessionAmbr"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	AddRedRanTunnelInfo             []QosFlowTunnel                 `json:"addRedRanTunnelInfo,omitempty"`
	SmfBindingInfo                  string                          `json:"smfBindingInfo,omitempty"`
	UdmGroupId                      string                          `json:"udmGroupId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	PduSessionType                  string                          `json:"pduSessionType"`
	PduSessionRef                   string                          `json:"pduSessionRef,omitempty"`
	SelMode                         string                          `json:"selMode,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	PduSessionId                    int                             `json:"pduSessionId"`
	MaxIntegrityProtectedDataRateDl string                          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	PduSessionSmfBinding            string                          `json:"pduSessionSmfBinding,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	PsaTunnelInfo                   *TunnelInfo                     `json:"psaTunnelInfo,omitempty"`
	SNssai                          Snssai                          `json:"sNssai"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
}