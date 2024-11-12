package models

type SmContextCreatedData struct {
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
}
