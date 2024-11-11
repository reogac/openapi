package models

type UeContextCreatedData struct {
	FailedSessionList  []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures  string            `json:"supportedFeatures,omitempty"`
	PcfReselectedInd   *bool             `json:"pcfReselectedInd,omitempty"`
	UeContext          UeContext         `json:"ueContext"`
	TargetToSourceData N2InfoContent     `json:"targetToSourceData"`
	PduSessionList     []N2SmInformation `json:"pduSessionList"`
}
