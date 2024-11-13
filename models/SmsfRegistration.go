package models

type SmsfRegistration struct {
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
}
