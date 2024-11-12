package models

type Amf3GppAccessRegistration struct {
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	Guami                       Guami                `json:"guami"`
	RatType                     RatType              `json:"ratType"`
	Pei                         string               `json:"pei,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
}
