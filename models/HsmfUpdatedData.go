package models
type HsmfUpdatedData struct {
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 QosFlowsSetupList	[]QosFlowSetupItem	`json:"qosFlowsSetupList,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 RoamingChargingProfile	*RoamingChargingProfile	`json:"roamingChargingProfile,omitempty"`
	 MaxIntegrityProtectedDataRateDl	string	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 MaxIntegrityProtectedDataRateUl	string	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 Ipv6MultiHomingInd	*bool	`json:"ipv6MultiHomingInd,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
}
