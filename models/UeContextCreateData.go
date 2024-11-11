package models

type UeContextCreateData struct {
	UeContext                  UeContext         `json:"ueContext"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
}
