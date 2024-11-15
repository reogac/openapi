/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistration struct {
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	Guami                       Guami           `json:"guami"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RatType                     RatType         `json:"ratType"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
}
