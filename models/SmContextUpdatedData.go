package models

type SmContextUpdatedData struct {
	Cause                    string                             `json:"cause,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N2SmInfoType             string                             `json:"n2SmInfoType,omitempty"`
	UpCnxState               string                             `json:"upCnxState,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	HoState                  string                             `json:"hoState,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
}
