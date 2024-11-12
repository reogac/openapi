package models

type HsmfUpdatedData struct {
	N1SmInfoToUe                    *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	N4InfoExt2                      *N4Information                `json:"n4InfoExt2,omitempty"`
	SessionAmbr                     *Ambr                         `json:"sessionAmbr,omitempty"`
	SupportedFeatures               string                        `json:"supportedFeatures,omitempty"`
	HomeProvidedChargingId          string                        `json:"homeProvidedChargingId,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem            `json:"qosFlowsSetupList,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	Pti                             *int                          `json:"pti,omitempty"`
	IntraPlmnApiRoot                string                        `json:"intraPlmnApiRoot,omitempty"`
	DnaiList                        []string                      `json:"dnaiList,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty"`
	Ipv6MultiHomingInd              *bool                         `json:"ipv6MultiHomingInd,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	InterPlmnApiRoot                string                        `json:"interPlmnApiRoot,omitempty"`
	N4Info                          *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt1                      *N4Information                `json:"n4InfoExt1,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
}
