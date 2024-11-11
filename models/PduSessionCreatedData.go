package models

type PduSessionCreatedData struct {
	EnablePauseCharging             *bool                           `json:"enablePauseCharging,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData                `json:"n1SmInfoToUe,omitempty"`
	AlwaysOnGranted                 *bool                           `json:"alwaysOnGranted,omitempty"`
	MaAcceptedInd                   *bool                           `json:"maAcceptedInd,omitempty"`
	PduSessionType                  string                          `json:"pduSessionType"`
	UeIpv4Address                   string                          `json:"ueIpv4Address,omitempty"`
	Gpsi                            string                          `json:"gpsi,omitempty"`
	Ipv6MultiHomingInd              *bool                           `json:"ipv6MultiHomingInd,omitempty"`
	NefExtBufSupportInd             *bool                           `json:"nefExtBufSupportInd,omitempty"`
	PduSessionId                    *int                            `json:"pduSessionId,omitempty"`
	UeIpv6Prefix                    string                          `json:"ueIpv6Prefix,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                 `json:"epsBearerInfo,omitempty"`
	SupportedFeatures               string                          `json:"supportedFeatures,omitempty"`
	DnaiList                        []string                        `json:"dnaiList,omitempty"`
	SmfServiceInstanceId            string                          `json:"smfServiceInstanceId,omitempty"`
	RedundantPduSessionInfo         *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem              `json:"qosFlowsSetupList,omitempty"`
	SmfInstanceId                   string                          `json:"smfInstanceId,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                  `json:"epsPdnCnxInfo,omitempty"`
	MaxIntegrityProtectedDataRate   string                          `json:"maxIntegrityProtectedDataRate,omitempty"`
	MaxIntegrityProtectedDataRateDl string                          `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	Ipv6Index                       *int                            `json:"ipv6Index,omitempty"`
	CnTunnelInfo                    *TunnelInfo                     `json:"cnTunnelInfo,omitempty"`
	AdditionalCnTunnelInfo          *TunnelInfo                     `json:"additionalCnTunnelInfo,omitempty"`
	SNssai                          *Snssai                         `json:"sNssai,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile         `json:"roamingChargingProfile,omitempty"`
	RecoveryTime                    string                          `json:"recoveryTime,omitempty"`
	SessionAmbr                     *Ambr                           `json:"sessionAmbr,omitempty"`
	HSmfServiceInstanceId           string                          `json:"hSmfServiceInstanceId,omitempty"`
	HomeProvidedChargingId          string                          `json:"homeProvidedChargingId,omitempty"`
	DnAaaAddress                    *IpAddress                      `json:"dnAaaAddress,omitempty"`
	UeIpv6InterfaceId               string                          `json:"ueIpv6InterfaceId,omitempty"`
	SscMode                         string                          `json:"sscMode"`
	HcnTunnelInfo                   *TunnelInfo                     `json:"hcnTunnelInfo,omitempty"`
	HSmfInstanceId                  string                          `json:"hSmfInstanceId,omitempty"`
	UpSecurity                      *UpSecurity                     `json:"upSecurity,omitempty"`
	SmallDataRateControlEnabled     *bool                           `json:"smallDataRateControlEnabled,omitempty"`
}
