package models

type HsmfUpdatedData struct {
	N4InfoExt2                      *N4Information                `json:"n4InfoExt2,omitempty"`
	SupportedFeatures               string                        `json:"supportedFeatures,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	InterPlmnApiRoot                string                        `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot                string                        `json:"intraPlmnApiRoot,omitempty"`
	Pti                             *int                          `json:"pti,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	DnaiList                        []string                      `json:"dnaiList,omitempty"`
	HomeProvidedChargingId          string                        `json:"homeProvidedChargingId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	Ipv6MultiHomingInd              *bool                         `json:"ipv6MultiHomingInd,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	N4Info                          *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt1                      *N4Information                `json:"n4InfoExt1,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SessionAmbr                     *Ambr                         `json:"sessionAmbr,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem            `json:"qosFlowsSetupList,omitempty"`
}
