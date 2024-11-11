package models

type UeContextCreateData struct {
	UeContext          UeContext         `json:"ueContext"`
	N2NotifyUri        string            `json:"n2NotifyUri,omitempty"`
	SupportedFeatures  string            `json:"supportedFeatures,omitempty"`
	UeRadioCapability  *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	NgapCause          *NgApCause        `json:"ngapCause,omitempty"`
	ServingNetwork     *PlmnIdNid        `json:"servingNetwork,omitempty"`
	TargetId           NgRanTargetId     `json:"targetId"`
	SourceToTargetData N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList     []N2SmInformation `json:"pduSessionList"`
}
