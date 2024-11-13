package models

type UeContextCreateData struct {
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
}
