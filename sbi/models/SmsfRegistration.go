/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
}
