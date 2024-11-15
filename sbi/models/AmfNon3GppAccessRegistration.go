/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:23 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RatType                     RatType         `json:"ratType"`
	Pei                         string          `json:"pei,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	Guami                       Guami           `json:"guami"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
}
