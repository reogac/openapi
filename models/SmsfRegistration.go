package models

type SmsfRegistration struct {
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
}
