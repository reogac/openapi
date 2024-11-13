package models

type AmfNon3GppAccessRegistration struct {
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	Guami                       Guami           `json:"guami"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	RatType                     RatType         `json:"ratType"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
}
