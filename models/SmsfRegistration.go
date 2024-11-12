package models

type SmsfRegistration struct {
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
}
