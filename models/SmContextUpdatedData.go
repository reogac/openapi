package models

type SmContextUpdatedData struct {
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
}
