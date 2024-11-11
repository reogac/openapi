package models

type UeContextTransferRspData struct {
	UeNbiotRadioCapability     *N2InfoContent `json:"ueNbiotRadioCapability,omitempty"`
	SupportedFeatures          string         `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext      `json:"ueContext"`
	UeRadioCapability          *N2InfoContent `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent `json:"ueRadioCapabilityForPaging,omitempty"`
}
