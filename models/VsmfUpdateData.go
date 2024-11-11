type VsmfUpdateData struct {
	 BackOffTimer	*int	`json:"backOffTimer,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 QosMonitoringInfo	*QosMonitoringInfo	`json:"qosMonitoringInfo,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 AdditionalCnTunnelInfo	*TunnelInfo	`json:"additionalCnTunnelInfo,omitempty"`
	 AssignEbiList	[]Arp	`json:"assignEbiList,omitempty"`
	 RevokeEbiList	[]int	`json:"revokeEbiList,omitempty"`
	 HsmfPduSessionUri	string	`json:"hsmfPduSessionUri,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 MaReleaseInd	string	`json:"maReleaseInd,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 QosFlowsAddModRequestList	[]QosFlowAddModifyRequestItem	`json:"qosFlowsAddModRequestList,omitempty"`
	 QosFlowsRelRequestList	[]QosFlowReleaseRequestItem	`json:"qosFlowsRelRequestList,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 SmallDataRateControlEnabled	*bool	`json:"smallDataRateControlEnabled,omitempty"`
	 RequestIndication	string	`json:"requestIndication"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 AlwaysOnGranted	*bool	`json:"alwaysOnGranted,omitempty"`
}
