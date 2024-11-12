package models

type UeContextCreateData struct {
	TargetId                   NgRanTargetId     `json:"targetId"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
}
