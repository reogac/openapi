package models
type AmfNon3GppAccessRegistration struct {
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 VgmlcAddress	*VgmlcAddress	`json:"vgmlcAddress,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 ImsVoPs	ImsVoPs	`json:"imsVoPs"`
	 UrrpIndicator	*bool	`json:"urrpIndicator,omitempty"`
	 NoEeSubscriptionInd	*bool	`json:"noEeSubscriptionInd,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 ReRegistrationRequired	*bool	`json:"reRegistrationRequired,omitempty"`
	 AdminDeregSubWithdrawn	*bool	`json:"adminDeregSubWithdrawn,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 SorSnpnSiSupported	*bool	`json:"sorSnpnSiSupported,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 AmfInstanceId	string	`json:"amfInstanceId"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 AmfServiceNamePcscfRest	ServiceName	`json:"amfServiceNamePcscfRest,omitempty"`
	 AmfServiceNameDereg	ServiceName	`json:"amfServiceNameDereg,omitempty"`
	 Guami	Guami	`json:"guami"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 BackupAmfInfo	[]BackupAmfInfo	`json:"backupAmfInfo,omitempty"`
	 RatType	RatType	`json:"ratType"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 PurgeFlag	*bool	`json:"purgeFlag,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri"`
	 AmfEeSubscriptionId	string	`json:"amfEeSubscriptionId,omitempty"`
}
