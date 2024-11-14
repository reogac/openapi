package models
type SmsfRegistration struct {
	 SmsfMAPAddress	string	`json:"smsfMAPAddress,omitempty"`
	 SmsfDiameterAddress	*NetworkNodeDiameterAddress	`json:"smsfDiameterAddress,omitempty"`
	 LastSynchronizationTime	string	`json:"lastSynchronizationTime,omitempty"`
	 ContextInfo	*ContextInfo	`json:"contextInfo,omitempty"`
	 ResetIds	[]string	`json:"resetIds,omitempty"`
	 UdrRestartInd	*bool	`json:"udrRestartInd,omitempty"`
	 SmsfSetId	string	`json:"smsfSetId,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SmsfSbiSupInd	*bool	`json:"smsfSbiSupInd,omitempty"`
	 DataRestorationCallbackUri	string	`json:"dataRestorationCallbackUri,omitempty"`
	 UeMemoryAvailableInd	*UeMemoryAvailableInd	`json:"ueMemoryAvailableInd,omitempty"`
	 SmsfInstanceId	string	`json:"smsfInstanceId"`
	 PlmnId	PlmnId	`json:"plmnId"`
	 RegistrationTime	string	`json:"registrationTime,omitempty"`
}
