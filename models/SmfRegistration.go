package models
type SmfRegistration struct {
	 PlmnId	PlmnId	`json:"plmnId"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 PduSessionId	int	`json:"pduSessionId"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 SmfInstanceId	string	`json:"smfInstanceId"`
	 SmfSetId	string	`json:"smfSetId,omitempty"`
	 PcfId	string	`json:"pcfId,omitempty"`
	 EpdgInd	*bool	`json:"epdgInd,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 SingleNssai	Snssai	`json:"singleNssai"`
	 PgwFqdn	string	`json:"pgwFqdn,omitempty"`
	 PgwIpAddr	*IpAddress	`json:"pgwIpAddr,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 EmergencyServices	*bool	`json:"emergencyServices,omitempty"`
	 RegistrationReason	RegistrationReason	`json:"registrationReason,omitempty"`
}
