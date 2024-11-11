package models

type HsmfUpdatedData struct {
	N4Info                          *N4Information          `json:"n4Info,omitempty"`
	N4InfoExt1                      *N4Information          `json:"n4InfoExt1,omitempty"`
	SupportedFeatures               string                  `json:"supportedFeatures,omitempty"`
	UpSecurity                      *UpSecurity             `json:"upSecurity,omitempty"`
	MaxIntegrityProtectedDataRateDl string                  `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData        `json:"n1SmInfoToUe,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
	HomeProvidedChargingId          string                  `json:"homeProvidedChargingId,omitempty"`
	Pti                             *int                    `json:"pti,omitempty"`
	N4InfoExt2                      *N4Information          `json:"n4InfoExt2,omitempty"`
	MaxIntegrityProtectedDataRateUl string                  `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	SessionAmbr                     *Ambr                   `json:"sessionAmbr,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo         `json:"epsBearerInfo,omitempty"`
	InterPlmnApiRoot                string                  `json:"interPlmnApiRoot,omitempty"`
	DnaiList                        []string                `json:"dnaiList,omitempty"`
	Ipv6MultiHomingInd              *bool                   `json:"ipv6MultiHomingInd,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem      `json:"qosFlowsSetupList,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo          `json:"epsPdnCnxInfo,omitempty"`
	IntraPlmnApiRoot                string                  `json:"intraPlmnApiRoot,omitempty"`
}
