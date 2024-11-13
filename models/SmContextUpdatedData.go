package models

type SmContextUpdatedData struct {
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
}
