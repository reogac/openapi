package models

type SmContextCreatedData struct {
	HoState              HoState          `json:"hoState,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
}
