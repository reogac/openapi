package models

type SmsfRegistration struct {
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
}
