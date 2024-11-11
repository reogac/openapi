package models

type SmContextUpdatedData struct {
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	UpCnxState               string                             `json:"upCnxState,omitempty"`
	HoState                  string                             `json:"hoState,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N2SmInfoType             string                             `json:"n2SmInfoType,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	Cause                    string                             `json:"cause,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
}
