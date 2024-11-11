package models

type SmContextCreatedData struct {
	Gpsi                 string           `json:"gpsi,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	UpCnxState           string           `json:"upCnxState,omitempty"`
	N2SmInfoType         string           `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	HoState              string           `json:"hoState,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
}
