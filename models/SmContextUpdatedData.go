package models

type SmContextUpdatedData struct {
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
}
