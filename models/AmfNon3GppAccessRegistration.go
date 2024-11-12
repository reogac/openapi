package models

type AmfNon3GppAccessRegistration struct {
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	Pei                         string          `json:"pei,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	RatType                     RatType         `json:"ratType"`
	Supi                        string          `json:"supi,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	Guami                       Guami           `json:"guami"`
}
