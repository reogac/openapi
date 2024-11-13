package models

type SmContextCreatedData struct {
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
}
