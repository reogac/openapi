package models

type UeContextCreateData struct {
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
}
