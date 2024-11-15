/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	Guami                       Guami                `json:"guami"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	RatType                     RatType              `json:"ratType"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
}
