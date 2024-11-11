package models

type SmContextCreatedData struct {
	UpCnxState           string           `json:"upCnxState,omitempty"`
	N2SmInfoType         string           `json:"n2SmInfoType,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	HoState              string           `json:"hoState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
}
