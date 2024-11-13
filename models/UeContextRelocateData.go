package models

type UeContextRelocateData struct {
	UeContext                UeContext         `json:"ueContext"`
	TargetId                 NgRanTargetId     `json:"targetId"`
	SourceToTargetData       N2InfoContent     `json:"sourceToTargetData"`
	ForwardRelocationRequest RefToBinaryData   `json:"forwardRelocationRequest"`
	PduSessionList           []N2SmInformation `json:"pduSessionList,omitempty"`
	UeRadioCapability        *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	NgapCause                *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures        string            `json:"supportedFeatures,omitempty"`
}
