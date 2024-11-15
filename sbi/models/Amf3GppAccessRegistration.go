/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami                `json:"guami"`
	RatType                     RatType              `json:"ratType"`
	Supi                        string               `json:"supi,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
}
