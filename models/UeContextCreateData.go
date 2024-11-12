package models

type UeContextCreateData struct {
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
}
