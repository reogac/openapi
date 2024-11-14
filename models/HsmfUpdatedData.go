package models
type HsmfUpdatedData struct {
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 HomeProvidedChargingId	string	`json:"homeProvidedChargingId,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 RoamingChargingProfile	*RoamingChargingProfile	`json:"roamingChargingProfile,omitempty"`
	 Ipv6MultiHomingInd	*bool	`json:"ipv6MultiHomingInd,omitempty"`
	 IntraPlmnApiRoot	string	`json:"intraPlmnApiRoot,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 MaxIntegrityProtectedDataRateDl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 MaxIntegrityProtectedDataRateUl	MaxIntegrityProtectedDataRate	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 QosFlowsSetupList	[]QosFlowSetupItem	`json:"qosFlowsSetupList,omitempty"`
}
