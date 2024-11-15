/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdatedData struct {
	DnaiList                        []string                      `json:"dnaiList,omitempty"`
	MaxIntegrityProtectedDataRateUl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	SessionAmbr                     *Ambr                         `json:"sessionAmbr,omitempty"`
	Pti                             *int                          `json:"pti,omitempty"`
	N1SmInfoToUe                    *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile       `json:"roamingChargingProfile,omitempty"`
	HomeProvidedChargingId          string                        `json:"homeProvidedChargingId,omitempty"`
	UpSecurity                      *UpSecurity                   `json:"upSecurity,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	InterPlmnApiRoot                string                        `json:"interPlmnApiRoot,omitempty"`
	IntraPlmnApiRoot                string                        `json:"intraPlmnApiRoot,omitempty"`
	N4Info                          *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt1                      *N4Information                `json:"n4InfoExt1,omitempty"`
	N4InfoExt2                      *N4Information                `json:"n4InfoExt2,omitempty"`
	SupportedFeatures               string                        `json:"supportedFeatures,omitempty"`
	Ipv6MultiHomingInd              *bool                         `json:"ipv6MultiHomingInd,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem            `json:"qosFlowsSetupList,omitempty"`
}
