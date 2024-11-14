/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	RatType                     RatType              `json:"ratType"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	Pei                         string               `json:"pei,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami                `json:"guami"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
}
