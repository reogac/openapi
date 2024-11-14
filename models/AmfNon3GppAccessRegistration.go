/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami           `json:"guami"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	RatType                     RatType         `json:"ratType"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
}
