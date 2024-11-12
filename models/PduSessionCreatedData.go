package models

type PduSessionCreatedData struct {
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	AdditionalSnssai                *Snssai                         `json:"additionalSnssai,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	SscMode                         string                          `json:"sscMode"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
}
