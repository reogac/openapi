/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	Guami                       Guami                `json:"guami"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	RatType                     RatType              `json:"ratType"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
}
