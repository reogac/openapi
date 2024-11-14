package models
type AmfNon3GppAccessRegistration struct {
	 Pei	string	`json:"pei,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 SorSnpnSiSupported	*bool	`json:"sorSnpnSiSupported,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 AmfInstanceId	string	`json:"amfInstanceId"`
	 VgmlcAddress	*VgmlcAddress	`json:"vgmlcAddress,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 RatType	RatType	`json:"ratType"`
	 PurgeFlag	*bool	`json:"purgeFlag,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 NoEeSubscriptionInd	*bool	`json:"noEeSubscriptionInd,omitempty"`
	 ReRegistrationRequired	*bool	`json:"reRegistrationRequired,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 BackupAmfInfo	[]BackupAmfInfo	`json:"backupAmfInfo,omitempty"`
	 AmfEeSubscriptionId	string	`json:"amfEeSubscriptionId,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 ImsVoPs	ImsVoPs	`json:"imsVoPs"`
	 UrrpIndicator	*bool	`json:"urrpIndicator,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 AdminDeregSubWithdrawn	*bool	`json:"adminDeregSubWithdrawn,omitempty"`
	 AmfServiceNameDereg	ServiceName	`json:"amfServiceNameDereg,omitempty"`
	 AmfServiceNamePcscfRest	ServiceName	`json:"amfServiceNamePcscfRest,omitempty"`
	 Guami	Guami	`json:"guami"`
}
