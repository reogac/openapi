package models
type SmfRegistration struct {
	 SingleNssai	Snssai	`json:"singleNssai"`
	 Dnn	string	`json:"dnn,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 PduSessionId	int	`json:"pduSessionId"`
	 EmergencyServices	*bool	`json:"emergencyServices,omitempty"`
	 PgwFqdn	string	`json:"pgwFqdn,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 PcfId	string	`json:"pcfId,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 SmfInstanceId	string	`json:"smfInstanceId"`
	 SmfSetId	string	`json:"smfSetId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 PgwIpAddr	*IpAddress	`json:"pgwIpAddr,omitempty"`
	 EpdgInd	*bool	`json:"epdgInd,omitempty"`
	 RegistrationReason	RegistrationReason	`json:"registrationReason,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
}
