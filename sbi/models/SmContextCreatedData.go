/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
}
