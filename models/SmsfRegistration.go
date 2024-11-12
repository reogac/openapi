package models

type SmsfRegistration struct {
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
}
