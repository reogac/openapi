package models

type HsmfUpdatedData struct {
	N4Info                          *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt1                      *N4Information                `json:"n4InfoExt1,omitempty"`
	N4InfoExt2                      *N4Information                `json:"n4InfoExt2,omitempty"`
	HomeProvidedChargingId          string                        `json:"homeProvidedChargingId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty"`
	Ipv6MultiHomingInd              *bool                         `json:"ipv6MultiHomingInd,omitempty"`
	IntraPlmnApiRoot                string                        `json:"intraPlmnApiRoot,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem            `json:"qosFlowsSetupList,omitempty"`
	SessionAmbr                     *Ambr                         `json:"sessionAmbr,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	DnaiList                        []string                      `json:"dnaiList,omitempty"`
	SupportedFeatures               string                        `json:"supportedFeatures,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	Pti                             *int                          `json:"pti,omitempty"`
	InterPlmnApiRoot                string                        `json:"interPlmnApiRoot,omitempty"`
}
