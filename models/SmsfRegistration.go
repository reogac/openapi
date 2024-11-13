package models

type SmsfRegistration struct {
	PlmnId                     PlmnId                      `json:"plmnId"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
}
