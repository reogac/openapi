package models

type PduSessionCreatedData struct {
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	PduSessionType                  string                          `json:"pduSessionType"`
	SscMode                         string                          `json:"sscMode"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	MaxIntegrityProtectedDataRateDl string                          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	MaxIntegrityProtectedDataRate   string                          `json:"maxIntegrityProtectedDataRate,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	AdditionalSnssai                *Snssai                         `json:"additionalSnssai,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
}
