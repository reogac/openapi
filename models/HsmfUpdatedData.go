type HsmfUpdatedData struct {
	 MaxIntegrityProtectedDataRateUl	string	`json:"maxIntegrityProtectedDataRateUl,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 UpSecurity	*UpSecurity	`json:"upSecurity,omitempty"`
	 Ipv6MultiHomingInd	*bool	`json:"ipv6MultiHomingInd,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 RoamingChargingProfile	*RoamingChargingProfile	`json:"roamingChargingProfile,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 QosFlowsSetupList	[]QosFlowSetupItem	`json:"qosFlowsSetupList,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 MaxIntegrityProtectedDataRateDl	string	`json:"maxIntegrityProtectedDataRateDl,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
}
