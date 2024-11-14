package models
type SmContextUpdatedData struct {
	 HoState	HoState	`json:"hoState,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 N3UlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3UlForwardingTnlList,omitempty"`
	 N9UlForwardingTunnel	*TunnelInfo	`json:"n9UlForwardingTunnel,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 ForwardingFTeid	string	`json:"forwardingFTeid,omitempty"`
	 AnchorSmfFeatures	*AnchorSmfFeatures	`json:"anchorSmfFeatures,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 N1SmMsg	*RefToBinaryData	`json:"n1SmMsg,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 EpsBearerSetup	[]string	`json:"epsBearerSetup,omitempty"`
	 DataForwarding	*bool	`json:"dataForwarding,omitempty"`
	 UpCnxState	UpCnxState	`json:"upCnxState,omitempty"`
	 ReleaseEbiList	[]int	`json:"releaseEbiList,omitempty"`
	 N3DlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3DlForwardingTnlList,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 ForwardingBearerContexts	[]string	`json:"forwardingBearerContexts,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
}
