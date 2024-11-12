package models

type SmContextCreatedData struct {
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
}
