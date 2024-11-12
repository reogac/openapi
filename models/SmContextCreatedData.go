package models

type SmContextCreatedData struct {
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
}
