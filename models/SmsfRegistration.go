/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
}
