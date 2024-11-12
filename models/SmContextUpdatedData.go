package models

type SmContextUpdatedData struct {
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
}
