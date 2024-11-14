package models
type Amf3GppAccessRegistration struct {
	 Pei	string	`json:"pei,omitempty"`
	 EpsInterworkingInfo	*EpsInterworkingInfo	`json:"epsInterworkingInfo,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri"`
	 Guami	Guami	`json:"guami"`
	 UrrpIndicator	*bool	`json:"urrpIndicator,omitempty"`
	 UeSrvccCapability	*bool	`json:"ueSrvccCapability,omitempty"`
	 ReRegistrationRequired	*bool	`json:"reRegistrationRequired,omitempty"`
	 AmfInstanceId	string	`json:"amfInstanceId"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 NoEeSubscriptionInd	*bool	`json:"noEeSubscriptionInd,omitempty"`
	 SorSnpnSiSupported	*bool	`json:"sorSnpnSiSupported,omitempty"`
	 EmergencyRegistrationInd	*bool	`json:"emergencyRegistrationInd,omitempty"`
	 RatType	RatType	`json:"ratType"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 AmfServiceNamePcscfRest	ServiceName	`json:"amfServiceNamePcscfRest,omitempty"`
	 VgmlcAddress	*VgmlcAddress	`json:"vgmlcAddress,omitempty"`
	 ImsVoPs	ImsVoPs	`json:"imsVoPs,omitempty"`
	 AmfServiceNameDereg	ServiceName	`json:"amfServiceNameDereg,omitempty"`
	 DrFlag	*bool	`json:"drFlag,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 UeReachableInd	UeReachableInd	`json:"ueReachableInd,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 PurgeFlag	*bool	`json:"purgeFlag,omitempty"`
	 BackupAmfInfo	[]BackupAmfInfo	`json:"backupAmfInfo,omitempty"`
	 AdminDeregSubWithdrawn	*bool	`json:"adminDeregSubWithdrawn,omitempty"`
	 UeMINTCapability	*bool	`json:"ueMINTCapability,omitempty"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 AmfEeSubscriptionId	string	`json:"amfEeSubscriptionId,omitempty"`
	 InitialRegistrationInd	*bool	`json:"initialRegistrationInd,omitempty"`
}
