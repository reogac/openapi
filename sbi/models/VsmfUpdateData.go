/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdateData struct {
	SessionAmbr                 *Ambr                         `json:"sessionAmbr,omitempty"`
	MaReleaseInd                MaReleaseIndication           `json:"maReleaseInd,omitempty"`
	SmallDataRateControlEnabled *bool                         `json:"smallDataRateControlEnabled,omitempty"`
	QosMonitoringInfo           *QosMonitoringInfo            `json:"qosMonitoringInfo,omitempty"`
	NewSmfId                    string                        `json:"newSmfId,omitempty"`
	N1smCause                   string                        `json:"n1smCause,omitempty"`
	AdditionalCnTunnelInfo      *TunnelInfo                   `json:"additionalCnTunnelInfo,omitempty"`
	RevokeEbiList               []int                         `json:"revokeEbiList,omitempty"`
	ModifiedEbiList             []EbiArpMapping               `json:"modifiedEbiList,omitempty"`
	AlwaysOnGranted             *bool                         `json:"alwaysOnGranted,omitempty"`
	NewSmfPduSessionUri         string                        `json:"newSmfPduSessionUri,omitempty"`
	N4InfoExt1                  *N4Information                `json:"n4InfoExt1,omitempty"`
	N9DataForwardingInd         *bool                         `json:"n9DataForwardingInd,omitempty"`
	QosFlowsAddModRequestList   []QosFlowAddModifyRequestItem `json:"qosFlowsAddModRequestList,omitempty"`
	BackOffTimer                *int                          `json:"backOffTimer,omitempty"`
	DnaiList                    []string                      `json:"dnaiList,omitempty"`
	N4InfoExt2                  *N4Information                `json:"n4InfoExt2,omitempty"`
	N4InfoExt3                  *N4Information                `json:"n4InfoExt3,omitempty"`
	RequestIndication           RequestIndication             `json:"requestIndication"`
	QosFlowsRelRequestList      []QosFlowReleaseRequestItem   `json:"qosFlowsRelRequestList,omitempty"`
	EpsBearerInfo               []EpsBearerInfo               `json:"epsBearerInfo,omitempty"`
	Cause                       Cause                         `json:"cause,omitempty"`
	N4Info                      *N4Information                `json:"n4Info,omitempty"`
	N9InactivityTimer           *int                          `json:"n9InactivityTimer,omitempty"`
	AssignEbiList               []Arp                         `json:"assignEbiList,omitempty"`
	Pti                         *int                          `json:"pti,omitempty"`
	N1SmInfoToUe                *RefToBinaryData              `json:"n1SmInfoToUe,omitempty"`
	HsmfPduSessionUri           string                        `json:"hsmfPduSessionUri,omitempty"`
	SupportedFeatures           string                        `json:"supportedFeatures,omitempty"`
	MaAcceptedInd               *bool                         `json:"maAcceptedInd,omitempty"`
	EpsPdnCnxInfo               *EpsPdnCnxInfo                `json:"epsPdnCnxInfo,omitempty"`
}