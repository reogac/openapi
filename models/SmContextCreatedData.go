package models

type SmContextCreatedData struct {
	SmfUri               string           `json:"smfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
}
