package models

type VsmfUpdateData struct {
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	RequestIndication           RequestIndication             `json:"requestIndication"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	Cause                       Cause                         `json:"cause,omitempty"`
	MaReleaseInd                MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
}
