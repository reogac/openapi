/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	RatType                     RatType         `json:"ratType"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	Guami                       Guami           `json:"guami"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
}
