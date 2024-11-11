package models

type VsmfUpdateData struct {
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	Cause                       string                        `json:"cause,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	RequestIndication           string                        `json:"requestIndication"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	MaReleaseInd                string                        `json:"maReleaseInd,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
}
