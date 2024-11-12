package models

type SmContextUpdatedData struct {
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
}
