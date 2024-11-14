package models
type SmContextUpdatedData struct {
	 N3UlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3UlForwardingTnlList,omitempty"`
	 Cause	Cause	`json:"cause,omitempty"`
	 AnchorSmfFeatures	*AnchorSmfFeatures	`json:"anchorSmfFeatures,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 EpsBearerSetup	[]string	`json:"epsBearerSetup,omitempty"`
	 N3DlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3DlForwardingTnlList,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 ReleaseEbiList	[]int	`json:"releaseEbiList,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 N1SmMsg	*RefToBinaryData	`json:"n1SmMsg,omitempty"`
	 DataForwarding	*bool	`json:"dataForwarding,omitempty"`
	 ForwardingBearerContexts	[]string	`json:"forwardingBearerContexts,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 UpCnxState	UpCnxState	`json:"upCnxState,omitempty"`
	 HoState	HoState	`json:"hoState,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 N9UlForwardingTunnel	*TunnelInfo	`json:"n9UlForwardingTunnel,omitempty"`
	 ForwardingFTeid	string	`json:"forwardingFTeid,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
}
