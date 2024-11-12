package models

type SmContextUpdatedData struct {
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	UpCnxState               string                             `json:"upCnxState,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfoType             string                             `json:"n2SmInfoType,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	HoState                  string                             `json:"hoState,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	Cause                    string                             `json:"cause,omitempty"`
}
