/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
}
