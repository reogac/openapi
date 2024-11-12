package models

type AmfNon3GppAccessRegistration struct {
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	RatType                     RatType         `json:"ratType"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	Guami                       Guami           `json:"guami"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
}
