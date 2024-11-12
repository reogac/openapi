package models

type AmfNon3GppAccessRegistration struct {
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	RatType                     RatType         `json:"ratType"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Guami                       Guami           `json:"guami"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
}
