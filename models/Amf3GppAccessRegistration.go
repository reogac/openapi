package models
type Amf3GppAccessRegistration struct {
	 UeReachableInd	UeReachableInd	`json:"ueReachableInd,omitempty"`
	 UeMINTCapability	*bool	`json:"ueMINTCapability,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 PcscfRestorationCallbackUri	string	`json:"pcscfRestorationCallbackUri,omitempty"`
	 RatType	RatType	`json:"ratType"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 DisasterRoamingInd	*bool	`json:"disasterRoamingInd,omitempty"`
	 DeregCallbackUri	string	`json:"deregCallbackUri"`
	 AmfServiceNameDereg	ServiceName	`json:"amfServiceNameDereg,omitempty"`
	 UrrpIndicator	*bool	`json:"urrpIndicator,omitempty"`
	 Supi	string	`json:"supi,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 InitialRegistrationInd	*bool	`json:"initialRegistrationInd,omitempty"`
	 UeSrvccCapability	*bool	`json:"ueSrvccCapability,omitempty"`
	 ReRegistrationRequired	*bool	`json:"reRegistrationRequired,omitempty"`
	 PurgeFlag	*bool	`json:"purgeFlag,omitempty"`
	 ImsVoPs	ImsVoPs	`json:"imsVoPs,omitempty"`
	 AmfServiceNamePcscfRest	ServiceName	`json:"amfServiceNamePcscfRest,omitempty"`
	 EmergencyRegistrationInd	*bool	`json:"emergencyRegistrationInd,omitempty"`
	 BackupAmfInfo	[]BackupAmfInfo	`json:"backupAmfInfo,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 NoEeSubscriptionInd	*bool	`json:"noEeSubscriptionInd,omitempty"`
	 AmfInstanceId	string	`json:"amfInstanceId"`
	 Guami	Guami	`json:"guami"`
	 VgmlcAddress	*VgmlcAddress	`json:"vgmlcAddress,omitempty"`
	 AdminDeregSubWithdrawn	*bool	`json:"adminDeregSubWithdrawn,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 AmfEeSubscriptionId	string	`json:"amfEeSubscriptionId,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 DrFlag	*bool	`json:"drFlag,omitempty"`
	 EpsInterworkingInfo	*EpsInterworkingInfo	`json:"epsInterworkingInfo,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 SorSnpnSiSupported	*bool	`json:"sorSnpnSiSupported,omitempty"`
}
