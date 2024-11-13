package models

type PduSessionCreatedData struct {
	PduSessionType                  PduSessionType                  `json:"pduSessionType"`
	SscMode                         string                          `json:"sscMode"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	NspuSupportInd                  *bool                           `json:"nspuSupportInd,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	InterPlmnApiRoot                string                          `json:"interPlmnApiRoot,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	AdditionalSnssai                *Snssai                         `json:"additionalSnssai,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	IntraPlmnApiRoot                string                          `json:"intraPlmnApiRoot,omitempty"`
	MaxIntegrityProtectedDataRate   MaxIntegrityProtectedDataRate   `json:"maxIntegrityProtectedDataRate,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
}
