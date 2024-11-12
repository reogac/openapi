package models

type SmContextCreatedData struct {
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              string           `json:"hoState,omitempty"`
	UpCnxState           string           `json:"upCnxState,omitempty"`
	N2SmInfoType         string           `json:"n2SmInfoType,omitempty"`
}
