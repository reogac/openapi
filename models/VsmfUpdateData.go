package models

type VsmfUpdateData struct {
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	MaReleaseInd                MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
	RequestIndication           RequestIndication             `json:"requestIndication"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	Cause                       Cause                         `json:"cause,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
}
