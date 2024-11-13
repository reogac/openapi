package models

type Amf3GppAccessRegistration struct {
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	Pei                         string               `json:"pei,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami                `json:"guami"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	RatType                     RatType              `json:"ratType"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
}
