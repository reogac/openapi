package models
type UeContextCreateData struct {
	 UeContext	UeContext	`json:"ueContext"`
	 SourceToTargetData	N2InfoContent	`json:"sourceToTargetData"`
	 PduSessionList	[]N2SmInformation	`json:"pduSessionList"`
	 N2NotifyUri	string	`json:"n2NotifyUri,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 TargetId	NgRanTargetId	`json:"targetId"`
	 UeRadioCapability	*N2InfoContent	`json:"ueRadioCapability,omitempty"`
	 UeRadioCapabilityForPaging	*N2InfoContent	`json:"ueRadioCapabilityForPaging,omitempty"`
	 NgapCause	*NgApCause	`json:"ngapCause,omitempty"`
	 ServingNetwork	*PlmnIdNid	`json:"servingNetwork,omitempty"`
}
