/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistration struct {
	Pei                         string               `json:"pei,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	Guami                       Guami                `json:"guami"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	RatType                     RatType              `json:"ratType"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
}
