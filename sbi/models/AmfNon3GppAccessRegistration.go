/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	AmfInstanceId               string          `json:"amfInstanceId"`
	Guami                       Guami           `json:"guami"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	RatType                     RatType         `json:"ratType"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
}
