package models

type VsmfUpdateData struct {
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	Cause                       Cause                         `json:"cause,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	RequestIndication           RequestIndication             `json:"requestIndication"`
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	MaReleaseInd                MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
}
