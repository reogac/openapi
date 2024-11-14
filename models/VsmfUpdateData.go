package models
type VsmfUpdateData struct {
	 AdditionalCnTunnelInfo	*TunnelInfo	`json:"additionalCnTunnelInfo,omitempty"`
	 AssignEbiList	[]Arp	`json:"assignEbiList,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N4InfoExt3	*N4Information	`json:"n4InfoExt3,omitempty"`
	 SmallDataRateControlEnabled	*bool	`json:"smallDataRateControlEnabled,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 N9InactivityTimer	*int	`json:"n9InactivityTimer,omitempty"`
	 BackOffTimer	*int	`json:"backOffTimer,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 QosMonitoringInfo	*QosMonitoringInfo	`json:"qosMonitoringInfo,omitempty"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 N9DataForwardingInd	*bool	`json:"n9DataForwardingInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 HsmfPduSessionUri	string	`json:"hsmfPduSessionUri,omitempty"`
	 NewSmfPduSessionUri	string	`json:"newSmfPduSessionUri,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 RequestIndication	RequestIndication	`json:"requestIndication"`
	 RevokeEbiList	[]int	`json:"revokeEbiList,omitempty"`
	 AlwaysOnGranted	*bool	`json:"alwaysOnGranted,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 MaReleaseInd	MaReleaseIndication	`json:"maReleaseInd,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 QosFlowsAddModRequestList	[]QosFlowAddModifyRequestItem	`json:"qosFlowsAddModRequestList,omitempty"`
	 QosFlowsRelRequestList	[]QosFlowReleaseRequestItem	`json:"qosFlowsRelRequestList,omitempty"`
}
