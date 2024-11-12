package models

type SmContextCreatedData struct {
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
}
