/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
}
