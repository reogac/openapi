package models

type AmfNon3GppAccessRegistration struct {
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami           `json:"guami"`
	Supi                        string          `json:"supi,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RatType                     RatType         `json:"ratType"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
}
