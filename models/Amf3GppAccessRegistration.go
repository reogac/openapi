package models

type Amf3GppAccessRegistration struct {
	RatType                     RatType              `json:"ratType"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	Guami                       Guami                `json:"guami"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
}
