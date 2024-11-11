type SmContextUpdatedData struct {
	 N1SmMsg	*RefToBinaryData	`json:"n1SmMsg,omitempty"`
	 N2SmInfoType	string	`json:"n2SmInfoType,omitempty"`
	 N3DlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3DlForwardingTnlList,omitempty"`
	 UpCnxState	string	`json:"upCnxState,omitempty"`
	 DataForwarding	*bool	`json:"dataForwarding,omitempty"`
	 Cause	string	`json:"cause,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 AnchorSmfFeatures	*AnchorSmfFeatures	`json:"anchorSmfFeatures,omitempty"`
	 HoState	string	`json:"hoState,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 ModifiedEbiList	[]EbiArpMapping	`json:"modifiedEbiList,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 EpsBearerSetup	[]string	`json:"epsBearerSetup,omitempty"`
	 N3UlForwardingTnlList	[]IndirectDataForwardingTunnelInfo	`json:"n3UlForwardingTnlList,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 ForwardingFTeid	string	`json:"forwardingFTeid,omitempty"`
	 ReleaseEbiList	[]int	`json:"releaseEbiList,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
	 ForwardingBearerContexts	[]string	`json:"forwardingBearerContexts,omitempty"`
}
