package models
type UeContextCreateData struct {
	 PduSessionList	[]N2SmInformation	`json:"pduSessionList"`
	 N2NotifyUri	string	`json:"n2NotifyUri,omitempty"`
	 UeRadioCapability	*N2InfoContent	`json:"ueRadioCapability,omitempty"`
	 UeRadioCapabilityForPaging	*N2InfoContent	`json:"ueRadioCapabilityForPaging,omitempty"`
	 ServingNetwork	*PlmnIdNid	`json:"servingNetwork,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 TargetId	NgRanTargetId	`json:"targetId"`
	 SourceToTargetData	N2InfoContent	`json:"sourceToTargetData"`
	 NgapCause	*NgApCause	`json:"ngapCause,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
}
