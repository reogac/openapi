package models
type SmsfRegistration struct {
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SmsfMAPAddress	string	`json:"smsfMAPAddress,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 SmsfDiameterAddress	*NetworkNodeDiameterAddress	`json:"smsfDiameterAddress,omitempty"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 UeMemoryAvailableInd	*UeMemoryAvailableInd	`json:"ueMemoryAvailableInd,omitempty"`
	 SmsfSbiSupInd	*bool	`json:"smsfSbiSupInd,omitempty"`
	 SmsfInstanceId	string	`json:"smsfInstanceId"`
	 SmsfSetId	string	`json:"smsfSetId,omitempty"`
	 PlmnId	PlmnId	`json:"plmnId"`
}
