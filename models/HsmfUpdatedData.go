package models

type HsmfUpdatedData struct {
	N4InfoExt1                      *N4Information          `json:"n4InfoExt1,omitempty"`
	SupportedFeatures               string                  `json:"supportedFeatures,omitempty"`
	Ipv6MultiHomingInd              *bool                   `json:"ipv6MultiHomingInd,omitempty"`
	InterPlmnApiRoot                string                  `json:"interPlmnApiRoot,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData        `json:"n1SmInfoToUe,omitempty"`
	DnaiList                        []string                `json:"dnaiList,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
	HomeProvidedChargingId          string                  `json:"homeProvidedChargingId,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem      `json:"qosFlowsSetupList,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo          `json:"epsPdnCnxInfo,omitempty"`
	SessionAmbr                     *Ambr                   `json:"sessionAmbr,omitempty"`
	Pti                             *int                    `json:"pti,omitempty"`
	IntraPlmnApiRoot                string                  `json:"intraPlmnApiRoot,omitempty"`
	N4Info                          *N4Information          `json:"n4Info,omitempty"`
	N4InfoExt2                      *N4Information          `json:"n4InfoExt2,omitempty"`
	UpSecurity                      *UpSecurity             `json:"upSecurity,omitempty"`
	MaxIntegrityProtectedDataRateUl string                  `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl string                  `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo         `json:"epsBearerInfo,omitempty"`
}
