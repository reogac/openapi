package models

type SmsfRegistration struct {
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
}
