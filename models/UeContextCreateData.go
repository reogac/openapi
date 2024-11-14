package models
type UeContextCreateData struct {
	 TargetId	NgRanTargetId	`json:"targetId"`
	 SourceToTargetData	N2InfoContent	`json:"sourceToTargetData"`
	 N2NotifyUri	string	`json:"n2NotifyUri,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 NgapCause	*NgApCause	`json:"ngapCause,omitempty"`
	 ServingNetwork	*PlmnIdNid	`json:"servingNetwork,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 PduSessionList	[]N2SmInformation	`json:"pduSessionList"`
	 UeRadioCapability	*N2InfoContent	`json:"ueRadioCapability,omitempty"`
	 UeRadioCapabilityForPaging	*N2InfoContent	`json:"ueRadioCapabilityForPaging,omitempty"`
}
