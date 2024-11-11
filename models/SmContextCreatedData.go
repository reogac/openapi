package models

type SmContextCreatedData struct {
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	HoState              string           `json:"hoState,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	UpCnxState           string           `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType         string           `json:"n2SmInfoType,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
}
