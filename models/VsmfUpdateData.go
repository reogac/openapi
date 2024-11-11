package models

type VsmfUpdateData struct {
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	MaReleaseInd                string                        `json:"maReleaseInd,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	Cause                       string                        `json:"cause,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
	RequestIndication           string                        `json:"requestIndication"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
}
