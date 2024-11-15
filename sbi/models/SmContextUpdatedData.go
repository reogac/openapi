/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextUpdatedData struct {
	ReleaseEbiList           []int                              `json:"releaseEbiList,omitempty"`
	N2SmInfo                 *RefToBinaryData                   `json:"n2SmInfo,omitempty"`
	N2SmInfoType             N2SmInfoType                       `json:"n2SmInfoType,omitempty"`
	DataForwarding           *bool                              `json:"dataForwarding,omitempty"`
	Cause                    Cause                              `json:"cause,omitempty"`
	SelectedSmfId            string                             `json:"selectedSmfId,omitempty"`
	AllocatedEbiList         []EbiArpMapping                    `json:"allocatedEbiList,omitempty"`
	N9UlForwardingTunnel     *TunnelInfo                        `json:"n9UlForwardingTunnel,omitempty"`
	ForwardingFTeid          string                             `json:"forwardingFTeid,omitempty"`
	SelectedOldSmfId         string                             `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot         string                             `json:"interPlmnApiRoot,omitempty"`
	EpsBearerSetup           []string                           `json:"epsBearerSetup,omitempty"`
	N3UlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3UlForwardingTnlList,omitempty"`
	MaAcceptedInd            *bool                              `json:"maAcceptedInd,omitempty"`
	ForwardingBearerContexts []string                           `json:"forwardingBearerContexts,omitempty"`
	AnchorSmfFeatures        *AnchorSmfFeatures                 `json:"anchorSmfFeatures,omitempty"`
	UpCnxState               UpCnxState                         `json:"upCnxState,omitempty"`
	HoState                  HoState                            `json:"hoState,omitempty"`
	ModifiedEbiList          []EbiArpMapping                    `json:"modifiedEbiList,omitempty"`
	N1SmMsg                  *RefToBinaryData                   `json:"n1SmMsg,omitempty"`
	N3DlForwardingTnlList    []IndirectDataForwardingTunnelInfo `json:"n3DlForwardingTnlList,omitempty"`
	SupportedFeatures        string                             `json:"supportedFeatures,omitempty"`
}
