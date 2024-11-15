/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
}
