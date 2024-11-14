package models
type HsmfUpdatedData struct {
	 HomeProvidedChargingId	string	`json:"homeProvidedChargingId,omitempty"`
	 MaxIntegrityProtectedDataRateUl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 MaxIntegrityProtectedDataRateDl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
	 IntraPlmnApiRoot	string	`json:"intraPlmnApiRoot,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 Ipv6MultiHomingInd	*bool	`json:"ipv6MultiHomingInd,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 RoamingChargingProfile	*RoamingChargingProfile	`json:"roamingChargingProfile,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 QosFlowsSetupList	[]QosFlowSetupItem	`json:"qosFlowsSetupList,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
}
