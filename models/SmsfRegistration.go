package models

type SmsfRegistration struct {
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
}
