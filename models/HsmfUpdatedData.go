package models

type HsmfUpdatedData struct {
	N1SmInfoToUe                    *RefToBinaryData        `json:"n1SmInfoToUe,omitempty"`
	N4InfoExt2                      *N4Information          `json:"n4InfoExt2,omitempty"`
	UpSecurity                      *UpSecurity             `json:"upSecurity,omitempty"`
	SessionAmbr                     *Ambr                   `json:"sessionAmbr,omitempty"`
	N4Info                          *N4Information          `json:"n4Info,omitempty"`
	HomeProvidedChargingId          string                  `json:"homeProvidedChargingId,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo          `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo         `json:"epsBearerInfo,omitempty"`
	N4InfoExt1                      *N4Information          `json:"n4InfoExt1,omitempty"`
	SupportedFeatures               string                  `json:"supportedFeatures,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
	MaxIntegrityProtectedDataRateDl string                  `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	Ipv6MultiHomingInd              *bool                   `json:"ipv6MultiHomingInd,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem      `json:"qosFlowsSetupList,omitempty"`
	InterPlmnApiRoot                string                  `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot                string                  `json:"intraPlmnApiRoot,omitempty"`
	DnaiList                        []string                `json:"dnaiList,omitempty"`
	MaxIntegrityProtectedDataRateUl string                  `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	Pti                             *int                    `json:"pti,omitempty"`
}
