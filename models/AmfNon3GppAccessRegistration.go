package models

type AmfNon3GppAccessRegistration struct {
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Guami                       Guami           `json:"guami"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	RatType                     RatType         `json:"ratType"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
}
