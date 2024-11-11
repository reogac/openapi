package models
type UeContextCreateData struct {
	 PduSessionList	[]N2SmInformation	`json:"pduSessionList"`
	 ServingNetwork	*PlmnIdNid	`json:"servingNetwork,omitempty"`
	 UeContext	UeContext	`json:"ueContext"`
	 SourceToTargetData	N2InfoContent	`json:"sourceToTargetData"`
	 UeRadioCapability	*N2InfoContent	`json:"ueRadioCapability,omitempty"`
	 NgapCause	*NgApCause	`json:"ngapCause,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 TargetId	NgRanTargetId	`json:"targetId"`
	 N2NotifyUri	string	`json:"n2NotifyUri,omitempty"`
}
