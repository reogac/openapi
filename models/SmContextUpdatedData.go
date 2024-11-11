package models

type SmContextUpdatedData struct {
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	Cause                    string                             `json:"cause,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	UpCnxState               string                             `json:"upCnxState,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N2SmInfoType             string                             `json:"n2SmInfoType,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	HoState                  string                             `json:"hoState,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
}
