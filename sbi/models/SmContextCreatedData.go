/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
}
