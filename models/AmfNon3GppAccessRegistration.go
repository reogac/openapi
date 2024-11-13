package models

type AmfNon3GppAccessRegistration struct {
	RegistrationTime            string          `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo    `json:"contextInfo,omitempty"`
	ResetIds                    []string        `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool           `json:"disasterRoamingInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName     `json:"amfServiceNamePcscfRest,omitempty"`
	VgmlcAddress                *VgmlcAddress   `json:"vgmlcAddress,omitempty"`
	AmfInstanceId               string          `json:"amfInstanceId"`
	PurgeFlag                   *bool           `json:"purgeFlag,omitempty"`
	BackupAmfInfo               []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string          `json:"amfEeSubscriptionId,omitempty"`
	DataRestorationCallbackUri  string          `json:"dataRestorationCallbackUri,omitempty"`
	ImsVoPs                     ImsVoPs         `json:"imsVoPs"`
	Pei                         string          `json:"pei,omitempty"`
	RatType                     RatType         `json:"ratType"`
	NoEeSubscriptionInd         *bool           `json:"noEeSubscriptionInd,omitempty"`
	ReRegistrationRequired      *bool           `json:"reRegistrationRequired,omitempty"`
	SupportedFeatures           string          `json:"supportedFeatures,omitempty"`
	AmfServiceNameDereg         ServiceName     `json:"amfServiceNameDereg,omitempty"`
	SorSnpnSiSupported          *bool           `json:"sorSnpnSiSupported,omitempty"`
	Guami                       Guami           `json:"guami"`
	LastSynchronizationTime     string          `json:"lastSynchronizationTime,omitempty"`
	DeregCallbackUri            string          `json:"deregCallbackUri"`
	UrrpIndicator               *bool           `json:"urrpIndicator,omitempty"`
	UdrRestartInd               *bool           `json:"udrRestartInd,omitempty"`
	PcscfRestorationCallbackUri string          `json:"pcscfRestorationCallbackUri,omitempty"`
	Supi                        string          `json:"supi,omitempty"`
	AdminDeregSubWithdrawn      *bool           `json:"adminDeregSubWithdrawn,omitempty"`
}
