package models

type SmContextUpdatedData struct {
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
}
