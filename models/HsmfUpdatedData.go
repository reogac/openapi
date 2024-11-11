package models

type HsmfUpdatedData struct {
	Pti                             *int                    `json:"pti,omitempty"`
	N4InfoExt1                      *N4Information          `json:"n4InfoExt1,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
	UpSecurity                      *UpSecurity             `json:"upSecurity,omitempty"`
	SessionAmbr                     *Ambr                   `json:"sessionAmbr,omitempty"`
	Ipv6MultiHomingInd              *bool                   `json:"ipv6MultiHomingInd,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo         `json:"epsBearerInfo,omitempty"`
	N4Info                          *N4Information          `json:"n4Info,omitempty"`
	DnaiList                        []string                `json:"dnaiList,omitempty"`
	MaxIntegrityProtectedDataRateUl string                  `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl string                  `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SupportedFeatures               string                  `json:"supportedFeatures,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem      `json:"qosFlowsSetupList,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData        `json:"n1SmInfoToUe,omitempty"`
	N4InfoExt2                      *N4Information          `json:"n4InfoExt2,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo          `json:"epsPdnCnxInfo,omitempty"`
}
