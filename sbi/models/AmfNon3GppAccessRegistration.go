/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	RatType                     RatType         `json:"ratType"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	Supi                        string          `json:"supi,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	Guami                       Guami           `json:"guami"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
}
