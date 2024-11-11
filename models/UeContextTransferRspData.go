package models

type UeContextTransferRspData struct {
	UeRadioCapability      *N2InfoContent `json:"ueRadioCapability,omitempty"`
	UeNbiotRadioCapability *N2InfoContent `json:"ueNbiotRadioCapability,omitempty"`
	SupportedFeatures      string         `json:"supportedFeatures,omitempty"`
	UeContext              UeContext      `json:"ueContext"`
}
