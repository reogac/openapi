package models

type PduSessionCreatedData struct {
	AdditionalSnssai                *Snssai                         `json:"additionalSnssai,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	MaxIntegrityProtectedDataRateDl string                          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	PduSessionType                  string                          `json:"pduSessionType"`
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	SscMode                         string                          `json:"sscMode"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	MaxIntegrityProtectedDataRate   string                          `json:"maxIntegrityProtectedDataRate,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
}
