package models

type UeContextCreateData struct {
	TargetId                   NgRanTargetId     `json:"targetId"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
}
