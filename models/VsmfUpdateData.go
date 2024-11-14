package models
type VsmfUpdateData struct {
	 QosFlowsAddModRequestList	[]QosFlowAddModifyRequestItem	`json:"qosFlowsAddModRequestList,omitempty"`
	 AssignEbiList	[]Arp	`json:"assignEbiList,omitempty"`
	 N1SmInfoToUe	*RefToBinaryData	`json:"n1SmInfoToUe,omitempty"`
	 HsmfPduSessionUri	string	`json:"hsmfPduSessionUri,omitempty"`
	 NewSmfPduSessionUri	string	`json:"newSmfPduSessionUri,omitempty"`
	 N4InfoExt3	*N4Information	`json:"n4InfoExt3,omitempty"`
	 SmallDataRateControlEnabled	*bool	`json:"smallDataRateControlEnabled,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 DnaiList	[]string	`json:"dnaiList,omitempty"`
	 RequestIndication	RequestIndication	`json:"requestIndication"`
	 NewSmfId	string	`json:"newSmfId,omitempty"`
	 BackOffTimer	*int	`json:"backOffTimer,omitempty"`
	 AdditionalCnTunnelInfo	*TunnelInfo	`json:"additionalCnTunnelInfo,omitempty"`
	 EpsPdnCnxInfo	*EpsPdnCnxInfo	`json:"epsPdnCnxInfo,omitempty"`
	 EpsBearerInfo	[]EpsBearerInfo	`json:"epsBearerInfo,omitempty"`
	 RevokeEbiList	[]int	`json:"revokeEbiList,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 AlwaysOnGranted	*bool	`json:"alwaysOnGranted,omitempty"`
	 MaReleaseInd	MaReleaseIndication	`json:"maReleaseInd,omitempty"`
	 N4InfoExt2	*N4Information	`json:"n4InfoExt2,omitempty"`
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 QosFlowsRelRequestList	[]QosFlowReleaseRequestItem	`json:"qosFlowsRelRequestList,omitempty"`
	 Pti	*int	`json:"pti,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 QosMonitoringInfo	*QosMonitoringInfo	`json:"qosMonitoringInfo,omitempty"`
	 N9InactivityTimer	*int	`json:"n9InactivityTimer,omitempty"`
	 N1smCause	string	`json:"n1smCause,omitempty"`
	 N4Info	*N4Information	`json:"n4Info,omitempty"`
	 N4InfoExt1	*N4Information	`json:"n4InfoExt1,omitempty"`
	 N9DataForwardingInd	*bool	`json:"n9DataForwardingInd,omitempty"`
}
