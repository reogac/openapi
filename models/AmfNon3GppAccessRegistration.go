package models

type AmfNon3GppAccessRegistration struct {
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RatType                     RatType         `json:"ratType"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami           `json:"guami"`
}
