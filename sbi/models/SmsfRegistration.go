/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
}
