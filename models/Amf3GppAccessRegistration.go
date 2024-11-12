package models

type Amf3GppAccessRegistration struct {
	ResetIds                    []string             `json:"resetIds,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	Guami                       Guami                `json:"guami"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	RatType                     RatType              `json:"ratType"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
}
